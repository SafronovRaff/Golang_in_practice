package main

import (
	"errors"
	"fmt"
	"os"
)

// label - уникальное наименование
type label string

// command - команда, которую можно выполнять в игре
type command label

// список доступных команд
var (
	eat  = command("eat")
	take = command("take")
	talk = command("talk to")
)

// thing - объект, который существует в игре
type thing struct {
	name    label
	actions map[command]string
}

// supports() возвращает true, если объект
// поддерживает команду action
func (t thing) supports(action command) bool {
	_, ok := t.actions[action]
	return ok
}

// String() возвращает описание объекта
func (t thing) String() string {
	return string(t.name)
}

// полный список объектов в игре
var (
	apple = thing{"apple", map[command]string{
		eat:  "mmm, delicious!",
		take: "you have an apple now",
	}}
	bob = thing{"bob", map[command]string{
		talk: "Bob says hello",
	}}
	coin = thing{"coin", map[command]string{
		take: "you have a coin now",
	}}
	mirror = thing{"mirror", map[command]string{
		take: "you have a mirror now",
		talk: "mirror does not answer",
	}}
	mushroom = thing{"mushroom", map[command]string{
		eat:  "tastes funny",
		take: "you have a mushroom now",
	}}
)

// step описывает шаг игры: сочетание команды и объекта
type step struct {
	cmd command
	obj thing
}

// isValid() возвращает true, если объект
// совместим с командой
func (s step) isValid() bool {
	return s.obj.supports(s.cmd)
}

// String() возвращает описание шага
func (s step) String() string {
	return fmt.Sprintf("%s %s", s.cmd, s.obj)
}

// //////////////////////////////////////////////// начало решения
// invalidStepError - ошибка, которая возникает,
// когда команда шага не совместима с объектом
type invalidStepError struct {
	Step step
}

func (er invalidStepError) Error() string {
	return fmt.Sprintf("cannot %s", er.Step)
}

// notEnoughObjectsError - ошибка, которая возникает,
// когда в игре закончились объекты определенного типа
type notEnoughObjectsError struct {
	object thing
}

func (er notEnoughObjectsError) Error() string {
	return fmt.Sprintf("there are no %ss left", er.object)
}

// commandLimitExceededError - ошибка, которая возникает,
// когда игрок превысил лимит на выполнение команды
type commandLimitExceededError struct {
	com command
}

func (er commandLimitExceededError) Error() string {
	return fmt.Sprintf("you don't want to %s anymore", er.com) //ты больше не хочешь ...
}

// objectLimitExceededError - ошибка, которая возникает,
// когда игрок превысил лимит на количество объектов
// определенного типа в инвентаре
type objectLimitExceededError struct {
	object thing
	limit  int
}

func (er objectLimitExceededError) Error() string {
	return fmt.Sprintf("you already have a %s", er.object) //у вас уже есть
}

// gameOverError - ошибка, которая произошла в игре
type gameOverError struct {
	// количество шагов, успешно выполненных
	// до того, как произошла ошибка
	nSteps int
	err    error
}

func (er gameOverError) Error() string {
	return er.err.Error()
}

func (er gameOverError) Unwrap() error {
	return er.err
}

// player - игрок
type player struct {
	// количество съеденного
	nEaten int
	// количество диалогов
	nDialogs int
	// инвентарь
	inventory []thing
}

// has() возвращает true, если у игрока
// в инвентаре есть предмет obj
func (p *player) has(obj thing) bool {
	for _, got := range p.inventory {
		if got.name == obj.name {
			return true
		}
	}
	return false
}

// do() выполняет команду cmd над объектом obj
// от имени игрока
func (p *player) do(cmd command, obj thing) error {
	// действуем в соответствии с командой
	switch cmd {
	case eat:
		if p.nEaten > 1 {
			return commandLimitExceededError{eat} //errors.New("you don't want to eat anymore") //ты больше не хочешь есть
		}
		p.nEaten++
	case take:
		if p.has(obj) {
			return objectLimitExceededError{object: obj, limit: 1} //fmt.Errorf("you already have a %s", obj) //у вас уже есть
		}
		p.inventory = append(p.inventory, obj)
	case talk:
		if p.nDialogs > 0 {
			return commandLimitExceededError{talk} //errors.New("you don't want to talk anymore") //ты больше не хочешь говорить
		}
		p.nDialogs++
	}
	return nil
}

// newPlayer создает нового игрока
func newPlayer() *player {
	return &player{inventory: []thing{}}
}

// game описывает игру
type game struct {
	// игрок
	player *player
	// объекты игрового мира
	things map[label]int
	// количество успешно выполненных шагов
	nSteps int
}

// has() проверяет, остались ли в игровом мире указанные предметы
func (g *game) has(obj thing) bool {
	count := g.things[obj.name]
	return count > 0
}

// execute() выполняет шаг step
func (g *game) execute(st step) error {
	// проверяем совместимость команды и объекта
	if !st.isValid() {
		return gameOverError{
			nSteps: g.nSteps,
			err:    invalidStepError{st},
		} //fmt.Errorf("cannot %s", st) //не могу
	}

	// когда игрок берет или съедает предмет,
	// тот пропадает из игрового мира
	if st.cmd == take || st.cmd == eat {
		if !g.has(st.obj) {
			return gameOverError{
				nSteps: g.nSteps,
				err:    notEnoughObjectsError{st.obj},
			}
			//fmt.Errorf("there are no %ss left", st.obj) //там не осталось
		}
		g.things[st.obj.name]--
	}

	// выполняем команду от имени игрока
	if err := g.player.do(st.cmd, st.obj); err != nil {
		return gameOverError{
			nSteps: g.nSteps,
			err:    err,
		}
	}

	g.nSteps++
	return nil
}

// newGame() создает новую игру
func newGame() *game {
	p := newPlayer()
	things := map[label]int{
		apple.name:    2,
		coin.name:     3,
		mirror.name:   1,
		mushroom.name: 1,
	}
	return &game{p, things, 0}
}

// giveAdvice() возвращает совет, который
// поможет игроку
//избежать ошибки err в будущем
/*Если команда не совместима с объектом, возвращает things like 'COMMAND OBJECT' are impossible, где COMMAND — название команды, а OBJECT — название объекта.

Если в игре закончились объекты определенного типа, возвращает be careful with scarce OBJECTs, где OBJECT — название объекта.

Если игрок слишком много ел, возвращает eat less. Если игрок слишком много говорил, возвращает talk to less.

Если игрок превысил лимит на количество объектов определенного типа в инвентаре, возвращает don't be greedy, LIMIT OBJECT is enough, где LIMIT — значение лимита, а OBJECT — название объекта.
Например:
things like 'eat bob' are impossible
be careful with scarce mirrors
don't be greedy, 1 apple is enough*/
func giveAdvice(err error) string {
	var invSrep invalidStepError
	if errors.As(err, &invSrep) {
		return fmt.Sprintf("things like '%v %v' are impossible", invSrep.Step.cmd, invSrep.Step.obj) // COMMAND — название команды, а OBJECT — название объекта.
	}
	var enOb notEnoughObjectsError
	if errors.As(err, &enOb) {
		return fmt.Sprintf("be careful with scarce %vs", enOb.object.name) //OBJECT — название объекта.
	}
	var comLim commandLimitExceededError
	if errors.As(err, &comLim) {
		if comLim.com == eat {
			return fmt.Sprintf("eat less")
		}
		if comLim.com == talk {
			return fmt.Sprintf("talk to less")
		}
		return ""
	}
	var obLimit objectLimitExceededError
	if errors.As(err, &obLimit) {
		return fmt.Sprintf("don't be greedy, %d %v is enough", obLimit.limit, obLimit.object.name) // где LIMIT — значение лимита, а OBJECT — название объекта
	}

	return ""
}

// конец решения

func main() {
	gm := newGame()
	steps := []step{
		{eat, apple},
		{talk, bob},
		{take, coin},
		{eat, mushroom},
	}

	for _, st := range steps {
		if err := tryStep(gm, st); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	fmt.Println("You win!")
}

// tryStep() выполняет шаг игры и печатает результат
func tryStep(gm *game, st step) error {
	fmt.Printf("trying to %s %s... ", st.cmd, st.obj.name)
	if err := gm.execute(st); err != nil {
		fmt.Println("FAIL")
		return err
	}
	fmt.Println("OK")
	return nil
}
