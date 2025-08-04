extends Node

@export var initState = Resource
const InitState = preload("res://scripts/resource/InitState.gd")

func _ready():
	var tech:int = initState.tech
	var res:int = initState.resource
	print(InitState.TechLevel.keys()[tech])
	print(InitState.ResourceLevel.keys()[res])
