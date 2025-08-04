enum TechLevel {
	Neolitic, IronAge, MedievalAge, Rennaissance, EarlyModern, Modern
}

enum ResourceLevel {
	None,
	Low,
	Standard,
	Plenty,
	Cheater
}

@export var tech: TechLevel = TechLevel.Neolitic
@export var resource: ResourceLevel = ResourceLevel.Low

func create():
	pass
