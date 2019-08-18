package logics

var HelloLogicHandler = HelloLogic{}

func init()  {
	HelloLogicHandler.LogicHandler = &HelloLogicHandler
}
