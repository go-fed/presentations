// C2S only
func NewSocialActor(c CommonBehavior,
	c2s SocialProtocol, // HLpro
	db Database,
	clock Clock) Actor

// S2S only
func NewFederatingActor(c CommonBehavior,
	s2s FederatingProtocol, // HLpro
	db Database,
	clock Clock) FederatingActor

// C2S & S2S
func NewActor(c CommonBehavior,
	c2s SocialProtocol, // HLpro
	s2s FederatingProtocol, // HLpro
	db Database,
	clock Clock) FederatingActor
