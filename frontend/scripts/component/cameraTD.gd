extends Camera3D

@export var speed = 400

var pin = false
var pin_mouse:Vector2

func GameCameraMove(delta: float):
	var axis:Vector2 = Vector2(0, 0)
	if Input.is_action_pressed("move_forward"):
		axis += Vector2(0, -1)
	if Input.is_action_pressed("move_backward"):
		axis += Vector2(0, 1)
	if Input.is_action_pressed("move_left"):
		axis += Vector2(-1, 0)
	if Input.is_action_pressed("move_right"):
		axis += Vector2(1, 0)
	axis = axis.normalized()
	
	translate_object_local(Vector3(
		axis.x * delta * speed, 
		0, 
		axis.y * delta * speed
	))

func GameCameraPin(delta: float):
	var pin = Input.is_mouse_button_pressed(MOUSE_BUTTON_RIGHT)
	if pin:
		var new_pin_mouse:Vector2 = get_viewport().get_mouse_position()
		var delta_pos = new_pin_mouse - pin_mouse
		pin_mouse = new_pin_mouse
		rotate_y(deg_to_rad(delta_pos.x))
		

func _ready():
	pass
	
func _process(delta: float):
	GameCameraMove(delta)
	GameCameraPin(delta)
