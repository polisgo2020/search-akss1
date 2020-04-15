package index

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"
)

func TestGetTokensFromText(t *testing.T) {
	content := `Prologue

The kids sat leaning over the side of the boat, looking into the water. "A huge pike," Anton said confidently.

"With fins this big?" Pashka asked.

Anton didn't reply. Anka also took a look, but saw only her own reflection.

"Be good to take a swim," said Pashka, plunging his arm into the water up to his elbow. "It's cold," he reported.`

	expectedJson := `{"anka":1,"anton":2,"arm":1,"asked":1,"big":1,"boat":1,"cold":1,"confidently":1,"didn":1,"elbow":1,"fins":1,"good":1,"huge":1,"kids":1,"leaning":1,"pashka":2,"pike":1,"plunging":1,"prologue":1,"reflection":1,"reply":1,"reported":1,"sat":1,"side":1,"swim":1,"water":2}`
	expected := Freq{}

	err := json.Unmarshal([]byte(expectedJson), &expected)
	if err != nil {
		t.Fatal(err)
	}

	actual := GetTokensFromText(content)
	if !reflect.DeepEqual(expected, actual) {
		log.Println("Tokens calculate failed")
		t.Fail()
	}
}

func TestReverseIdx_Search(t *testing.T) {
	idxMap := `{"absentmindedly":[{"prisoners":1}],"absolutely":[{"prisoners":1}],"absurd":[{"prisoners":1}],"absurdly":[{"prisoners":1}],"accidents":[{"prisoners":1}],"acquiring":[{"prisoners":1}],"activated":[{"prisoners":1}],"adventure":[{"prisoners":1}],"adventures":[{"prisoners":1}],"afar":[{"prisoners":1}],"affairs":[{"prisoners":1}],"age":[{"prisoners":1}],"ago":[{"prisoners":1}],"ahead":[{"prisoners":1}],"air":[{"prisoners":3}],"airy":[{"prisoners":1}],"alien":[{"prisoners":2}],"alleys":[{"prisoners":1}],"analyses":[{"prisoners":1}],"ancient":[{"prisoners":1}],"animal":[{"prisoners":2}],"anka":[{"hard":3}],"anton":[{"hard":6}],"ape":[{"prisoners":1}],"appealing":[{"prisoners":1}],"appeared":[{"prisoners":3}],"appetite":[{"prisoners":1}],"appetizing":[{"prisoners":1}],"approaching":[{"prisoners":1}],"aptitude":[{"hard":1}],"arbalest":[{"hard":1}],"area":[{"prisoners":3}],"arm":[{"hard":1}],"asked":[{"hard":1}],"assembled":[{"prisoners":1}],"assistant":[{"prisoners":1}],"atlas":[{"prisoners":1}],"attack":[{"prisoners":1}],"attacks":[{"prisoners":1}],"attract":[{"prisoners":1}],"automatic":[{"prisoners":1}],"awake":[{"hard":1}],"awe":[{"prisoners":1}],"awful":[{"prisoners":1}],"backed":[{"prisoners":1}],"bam":[{"prisoners":1}],"bank":[{"prisoners":4}],"barbarous":[{"prisoners":1}],"bare":[{"prisoners":1}],"barely":[{"prisoners":1}],"batteries":[{"prisoners":1}],"battle":[{"prisoners":1}],"beams":[{"prisoners":1}],"bear":[{"prisoners":1}],"beasts":[{"prisoners":1}],"beauties":[{"prisoners":1}],"began":[{"prisoners":1}],"beings":[{"prisoners":1}],"beneath":[{"prisoners":3}],"bet":[{"prisoners":1}],"better":[{"hard":1},{"prisoners":2}],"biblical":[{"prisoners":1}],"big":[{"hard":1},{"prisoners":1}],"bizarre":[{"prisoners":1}],"black":[{"hard":2}],"blindfolded":[{"prisoners":1}],"blindly":[{"prisoners":1}],"blots":[{"prisoners":2}],"blue":[{"hard":3},{"prisoners":2}],"blueness":[{"prisoners":1}],"boat":[{"hard":3}],"bodies":[{"prisoners":1}],"body":[{"hard":1},{"prisoners":1}],"bog":[{"prisoners":2}],"bolder":[{"prisoners":1}],"bolt":[{"prisoners":1}],"bottom":[{"prisoners":1}],"boulder":[{"prisoners":1}],"boulders":[{"prisoners":1}],"bound":[{"prisoners":1}],"boy":[{"prisoners":2}],"branches":[{"prisoners":1}],"breaking":[{"prisoners":1}],"breath":[{"prisoners":1}],"bridge":[{"prisoners":1}],"bright":[{"prisoners":1}],"brilliant":[{"prisoners":1}],"brittle":[{"prisoners":1}],"broad":[{"prisoners":2}],"broke":[{"prisoners":3}],"brushed":[{"prisoners":1}],"buckled":[{"prisoners":1}],"buddy":[{"prisoners":1}],"builders":[{"prisoners":1}],"buildings":[{"prisoners":1}],"bulrushes":[{"prisoners":2}],"burned":[{"prisoners":1}],"burst":[{"prisoners":1}],"bushes":[{"prisoners":4}],"buzzed":[{"prisoners":1}],"cabin":[{"prisoners":1}],"cable":[{"hard":1}],"carbon":[{"prisoners":1}],"careful":[{"hard":1}],"carefully":[{"prisoners":2}],"carrying":[{"prisoners":1}],"castle":[{"prisoners":1}],"castles":[{"prisoners":1}],"catalog":[{"prisoners":2}],"caterpillar":[{"prisoners":1}],"caught":[{"prisoners":2}],"cautiously":[{"prisoners":1}],"ceramic":[{"prisoners":1}],"chanced":[{"prisoners":1}],"change":[{"prisoners":1}],"characteristics":[{"prisoners":1}],"check":[{"prisoners":1}],"cheeks":[{"prisoners":1}],"childish":[{"hard":1}],"choked":[{"prisoners":1}],"choose":[{"prisoners":1}],"chrome":[{"hard":1}],"circle":[{"prisoners":2}],"circuit":[{"prisoners":1}],"city":[{"prisoners":1}],"civilization":[{"prisoners":2}],"civilizations":[{"prisoners":1}],"clambered":[{"hard":1}],"clanging":[{"prisoners":2}],"clear":[{"hard":1},{"prisoners":3}],"clearly":[{"prisoners":1}],"cliff":[{"hard":1}],"clogged":[{"prisoners":1}],"closer":[{"prisoners":4}],"clouds":[{"prisoners":1}],"clumsy":[{"prisoners":1}],"coarse":[{"prisoners":1}],"coil":[{"prisoners":4}],"cold":[{"hard":1},{"prisoners":1}],"collars":[{"prisoners":1}],"column":[{"prisoners":1}],"columns":[{"prisoners":1}],"concrete":[{"prisoners":4}],"confidently":[{"hard":1}],"consciousness":[{"prisoners":1}],"construct":[{"prisoners":1}],"construction":[{"prisoners":1}],"contact":[{"prisoners":2}],"contemplated":[{"prisoners":1}],"continued":[{"prisoners":1}],"contribution":[{"prisoners":1}],"controls":[{"prisoners":2}],"convinced":[{"prisoners":1}],"cool":[{"prisoners":1}],"copper":[{"hard":1}],"corkscrew":[{"prisoners":1}],"corridor":[{"prisoners":1}],"council":[{"prisoners":1}],"course":[{"prisoners":2}],"cracked":[{"prisoners":1}],"crackled":[{"prisoners":1}],"crackling":[{"prisoners":2}],"crash":[{"prisoners":1}],"crashing":[{"prisoners":1}],"creating":[{"prisoners":1}],"creatures":[{"prisoners":1}],"cries":[{"prisoners":1}],"crossbow":[{"hard":1}],"crossbows":[{"hard":1}],"crossroad":[{"prisoners":1}],"crusoe":[{"prisoners":3}],"culture":[{"prisoners":1}],"cut":[{"prisoners":2}],"cutting":[{"prisoners":1}],"cyberpilot":[{"prisoners":1}],"damn":[{"prisoners":1}],"damned":[{"prisoners":1}],"damp":[{"prisoners":1}],"danced":[{"prisoners":1}],"dangerously":[{"prisoners":1}],"dangling":[{"prisoners":1}],"dark":[{"hard":1},{"prisoners":4}],"darkly":[{"prisoners":1}],"darted":[{"prisoners":1}],"dawned":[{"prisoners":1}],"day":[{"prisoners":1}],"days":[{"prisoners":1}],"dead":[{"prisoners":1}],"deal":[{"prisoners":1}],"death":[{"prisoners":3}],"deck":[{"prisoners":1}],"decrepit":[{"prisoners":1}],"deep":[{"hard":1},{"prisoners":2}],"deeper":[{"prisoners":1}],"deeply":[{"prisoners":1}],"deer":[{"prisoners":1}],"dense":[{"prisoners":4}],"denser":[{"prisoners":1}],"dent":[{"prisoners":1}],"depths":[{"prisoners":2}],"descending":[{"prisoners":1}],"deserted":[{"prisoners":1}],"desolate":[{"prisoners":1}],"determine":[{"prisoners":1}],"device":[{"prisoners":1}],"didn":[{"hard":2},{"prisoners":1}],"die":[{"prisoners":1}],"died":[{"prisoners":1}],"difficult":[{"prisoners":1}],"dim":[{"prisoners":1}],"dioxide":[{"prisoners":1}],"dirty":[{"prisoners":1}],"disappeared":[{"prisoners":3}],"discover":[{"prisoners":1}],"discoveries":[{"prisoners":1}],"discussions":[{"prisoners":1}],"disgraced":[{"prisoners":1}],"dispatched":[{"prisoners":1}],"distressed":[{"prisoners":1}],"doesn":[{"prisoners":1}],"don":[{"prisoners":3}],"downhill":[{"prisoners":1}],"dragon":[{"prisoners":2}],"drawing":[{"prisoners":1}],"dreary":[{"prisoners":1}],"dressed":[{"prisoners":1}],"drew":[{"prisoners":2}],"dried":[{"prisoners":1}],"driverless":[{"prisoners":1}],"dropped":[{"prisoners":1}],"dry":[{"prisoners":2}],"dust":[{"prisoners":1}],"dusty":[{"prisoners":1}],"earth":[{"prisoners":4}],"east":[{"prisoners":2}],"easy":[{"prisoners":1}],"echoed":[{"prisoners":1}],"edge":[{"prisoners":1}],"edged":[{"prisoners":1}],"edging":[{"prisoners":1}],"edible":[{"prisoners":2}],"elbow":[{"hard":1}],"emerged":[{"prisoners":1}],"emitting":[{"prisoners":1}],"emptying":[{"prisoners":1}],"encased":[{"prisoners":1}],"enchanted":[{"prisoners":1}],"encounter":[{"prisoners":1}],"encountered":[{"prisoners":1}],"endless":[{"prisoners":1}],"engine":[{"prisoners":1}],"enormous":[{"prisoners":3}],"enthusiastic":[{"prisoners":1}],"entire":[{"prisoners":2}],"enveloped":[{"prisoners":1}],"episode":[{"prisoners":1}],"equator":[{"prisoners":1}],"escape":[{"prisoners":1}],"establish":[{"prisoners":1}],"establishing":[{"prisoners":1}],"establishment":[{"prisoners":1}],"eventually":[{"prisoners":1}],"evil":[{"prisoners":1}],"exhaust":[{"prisoners":1}],"exhausting":[{"prisoners":2}],"exhaustingly":[{"prisoners":1}],"expectantly":[{"hard":1}],"expected":[{"prisoners":2}],"experiences":[{"prisoners":1}],"expert":[{"prisoners":1}],"experts":[{"prisoners":1}],"exploded":[{"prisoners":1}],"exploration":[{"prisoners":1}],"expression":[{"prisoners":1}],"eye":[{"prisoners":1}],"fabulous":[{"prisoners":1}],"face":[{"prisoners":3}],"factories":[{"prisoners":1}],"failing":[{"prisoners":1}],"fair":[{"prisoners":1}],"fairy":[{"prisoners":2}],"fantastic":[{"prisoners":1}],"fashioned":[{"hard":1}],"father":[{"prisoners":2}],"feed":[{"prisoners":1}],"feet":[{"prisoners":4}],"felt":[{"prisoners":2}],"field":[{"prisoners":1}],"fighting":[{"prisoners":1}],"figure":[{"prisoners":1}],"filled":[{"prisoners":1}],"find":[{"prisoners":5}],"fingers":[{"prisoners":1}],"finished":[{"prisoners":1}],"fins":[{"hard":1}],"fire":[{"prisoners":4}],"firmament":[{"prisoners":3}],"flames":[{"prisoners":1}],"flared":[{"prisoners":1}],"flash":[{"prisoners":1}],"flat":[{"prisoners":1}],"flesh":[{"prisoners":1}],"flinging":[{"prisoners":1}],"floated":[{"prisoners":1}],"flow":[{"prisoners":1}],"flowed":[{"prisoners":1}],"fog":[{"hard":1}],"food":[{"prisoners":1}],"foot":[{"prisoners":1}],"footprints":[{"prisoners":1}],"forced":[{"prisoners":2}],"foresee":[{"prisoners":1}],"forest":[{"hard":1},{"prisoners":12}],"forever":[{"prisoners":1}],"forward":[{"prisoners":1}],"foul":[{"prisoners":1}],"fragments":[{"prisoners":1}],"fried":[{"hard":1}],"front":[{"hard":1},{"prisoners":2}],"fumes":[{"prisoners":1}],"fury":[{"prisoners":1}],"galactic":[{"prisoners":1}],"game":[{"prisoners":2}],"gaze":[{"prisoners":1}],"giant":[{"hard":1},{"prisoners":1}],"girders":[{"prisoners":1}],"glow":[{"prisoners":1}],"glowed":[{"prisoners":3}],"gnarled":[{"hard":1},{"prisoners":3}],"god":[{"prisoners":1}],"good":[{"hard":1}],"grab":[{"prisoners":1}],"grabbed":[{"hard":1}],"gradually":[{"prisoners":2}],"grass":[{"prisoners":6}],"gray":[{"prisoners":1}],"great":[{"prisoners":2}],"green":[{"hard":2}],"grew":[{"prisoners":2}],"grinding":[{"prisoners":2}],"grizzled":[{"hard":1}],"growing":[{"prisoners":1}],"growling":[{"prisoners":1}],"guilty":[{"prisoners":1}],"guttural":[{"prisoners":1}],"habitation":[{"prisoners":1}],"half":[{"prisoners":2}],"hand":[{"prisoners":1}],"handful":[{"prisoners":1}],"hands":[{"prisoners":1}],"happen":[{"prisoners":1}],"happened":[{"prisoners":3}],"happier":[{"prisoners":1}],"happy":[{"prisoners":1}],"hard":[{"prisoners":1}],"hatch":[{"prisoners":3}],"haunches":[{"prisoners":1}],"haven":[{"prisoners":5}],"hazy":[{"prisoners":1}],"head":[{"prisoners":2}],"headed":[{"prisoners":1}],"headquarters":[{"prisoners":1}],"hearing":[{"prisoners":1}],"heat":[{"prisoners":2}],"heavily":[{"prisoners":1}],"heavy":[{"prisoners":1}],"help":[{"prisoners":1}],"helpless":[{"prisoners":1}],"hesitatingly":[{"prisoners":1}],"high":[{"prisoners":3}],"highway":[{"prisoners":1}],"hippo":[{"prisoners":1}],"hobbled":[{"prisoners":1}],"hold":[{"prisoners":1}],"hole":[{"prisoners":2}],"hollow":[{"prisoners":1}],"hoof":[{"prisoners":1}],"hope":[{"prisoners":1}],"horizon":[{"prisoners":1}],"hot":[{"prisoners":3}],"hours":[{"prisoners":1}],"houses":[{"prisoners":1}],"huge":[{"hard":1},{"prisoners":1}],"human":[{"prisoners":3}],"humanoid":[{"prisoners":1}],"humpbacked":[{"prisoners":1}],"hunching":[{"prisoners":1}],"hunger":[{"prisoners":2}],"hunt":[{"prisoners":1}],"hunting":[{"prisoners":1}],"idea":[{"prisoners":1}],"imagine":[{"prisoners":1}],"imagined":[{"prisoners":1}],"impact":[{"prisoners":1}],"impossible":[{"prisoners":1}],"impudent":[{"prisoners":1}],"inability":[{"prisoners":1}],"incomprehensible":[{"prisoners":1}],"incredible":[{"prisoners":3}],"independent":[{"prisoners":1}],"infinite":[{"prisoners":1}],"inhabitants":[{"prisoners":1}],"inhabited":[{"prisoners":6}],"inhaling":[{"prisoners":1}],"inside":[{"prisoners":1}],"insulted":[{"prisoners":1}],"intelligent":[{"prisoners":2}],"intense":[{"prisoners":1}],"intention":[{"prisoners":1}],"interesting":[{"prisoners":1}],"interfered":[{"prisoners":1}],"interlaced":[{"prisoners":1}],"intersection":[{"prisoners":4}],"invisible":[{"prisoners":1}],"iron":[{"prisoners":6}],"iru":[{"prisoners":4}],"island":[{"prisoners":2}],"jagged":[{"prisoners":1}],"jerked":[{"hard":1}],"job":[{"prisoners":1}],"jolted":[{"prisoners":1}],"jumped":[{"hard":1},{"prisoners":2}],"junk":[{"prisoners":1}],"jutted":[{"hard":1}],"jutting":[{"prisoners":1}],"kids":[{"hard":1},{"prisoners":1}],"kind":[{"prisoners":2}],"king":[{"hard":1}],"knew":[{"prisoners":1}],"knocked":[{"prisoners":1}],"lab":[{"prisoners":1}],"laboratory":[{"prisoners":1}],"lacked":[{"hard":1},{"prisoners":1}],"lad":[{"prisoners":1}],"lake":[{"hard":2}],"landed":[{"hard":1}],"landing":[{"prisoners":1}],"landscape":[{"prisoners":2}],"lanthanides":[{"prisoners":1}],"large":[{"prisoners":2}],"latticed":[{"prisoners":1}],"lay":[{"prisoners":1}],"lazy":[{"hard":1}],"lead":[{"prisoners":1}],"leading":[{"prisoners":1}],"leaned":[{"prisoners":1}],"leaning":[{"hard":1}],"leaped":[{"prisoners":1}],"learned":[{"prisoners":1}],"led":[{"prisoners":1}],"left":[{"prisoners":2}],"lengthening":[{"prisoners":1}],"lengthy":[{"prisoners":1}],"lever":[{"hard":1}],"life":[{"prisoners":2}],"lifted":[{"prisoners":1}],"light":[{"prisoners":1}],"lightning":[{"prisoners":1}],"listen":[{"prisoners":1}],"listened":[{"prisoners":2}],"lit":[{"prisoners":1}],"litter":[{"prisoners":2}],"ll":[{"prisoners":2}],"local":[{"prisoners":2}],"locating":[{"prisoners":1}],"logical":[{"prisoners":1}],"long":[{"prisoners":4}],"looked":[{"hard":2},{"prisoners":4}],"loomed":[{"prisoners":1}],"lost":[{"prisoners":1}],"loud":[{"prisoners":1}],"lousy":[{"prisoners":1}],"low":[{"prisoners":1}],"lower":[{"hard":1}],"luck":[{"prisoners":1}],"luminous":[{"prisoners":1}],"lungs":[{"prisoners":1}],"lying":[{"prisoners":2}],"machine":[{"prisoners":2}],"main":[{"prisoners":2}],"making":[{"prisoners":1}],"mammoth":[{"prisoners":2}],"manage":[{"prisoners":1}],"marshal":[{"hard":1}],"material":[{"prisoners":2}],"maxie":[{"prisoners":1}],"maxim":[{"prisoners":20}],"meat":[{"prisoners":1}],"mechanical":[{"hard":1},{"prisoners":1}],"meet":[{"prisoners":1}],"menacing":[{"prisoners":2}],"mentally":[{"prisoners":1}],"mess":[{"prisoners":1}],"metal":[{"prisoners":2}],"meteorite":[{"prisoners":2}],"meteorites":[{"prisoners":1}],"midges":[{"prisoners":3}],"mile":[{"prisoners":1}],"mind":[{"prisoners":4}],"minute":[{"prisoners":1}],"missed":[{"prisoners":1}],"mistake":[{"prisoners":1}],"mixed":[{"prisoners":1}],"mixture":[{"prisoners":1}],"mm":[{"prisoners":1}],"moist":[{"prisoners":1}],"monotonous":[{"prisoners":3}],"monotonously":[{"prisoners":2}],"monster":[{"prisoners":2}],"monsters":[{"prisoners":1}],"moon":[{"prisoners":1}],"moonlit":[{"prisoners":1}],"mosquitoes":[{"prisoners":1}],"mother":[{"prisoners":3}],"motion":[{"hard":1}],"moved":[{"prisoners":1}],"moving":[{"prisoners":2}],"mud":[{"prisoners":1}],"muddy":[{"prisoners":1}],"muffled":[{"prisoners":1}],"multitude":[{"prisoners":1}],"mumbled":[{"prisoners":1}],"musty":[{"prisoners":1}],"naive":[{"prisoners":1}],"narrow":[{"prisoners":1}],"nearby":[{"prisoners":2}],"neck":[{"hard":1}],"necks":[{"prisoners":1}],"newfangled":[{"hard":1}],"nib":[{"prisoners":1}],"night":[{"prisoners":1}],"nocturnal":[{"prisoners":1}],"noiselessly":[{"hard":1},{"prisoners":1}],"noises":[{"prisoners":1}],"normal":[{"prisoners":1}],"north":[{"hard":1}],"noticed":[{"prisoners":2}],"notion":[{"prisoners":1}],"oar":[{"hard":1}],"obstacles":[{"prisoners":1}],"occasionally":[{"prisoners":1}],"occurred":[{"prisoners":1}],"odor":[{"prisoners":1}],"offer":[{"prisoners":1}],"open":[{"prisoners":4}],"opened":[{"prisoners":1}],"operated":[{"hard":1}],"opposite":[{"prisoners":2}],"orbit":[{"prisoners":1}],"outgrow":[{"prisoners":1}],"outgrown":[{"prisoners":1}],"overgrown":[{"prisoners":2}],"overhead":[{"prisoners":1}],"overlaid":[{"hard":1}],"overtake":[{"prisoners":2}],"ox":[{"hard":1}],"paces":[{"prisoners":1}],"pair":[{"prisoners":1}],"pal":[{"hard":1}],"pale":[{"hard":1}],"palm":[{"prisoners":2}],"panic":[{"prisoners":1}],"parts":[{"prisoners":1}],"pashka":[{"hard":6}],"pass":[{"prisoners":1}],"passed":[{"prisoners":1}],"paused":[{"prisoners":1}],"pebbles":[{"prisoners":1}],"people":[{"prisoners":2}],"perform":[{"prisoners":1}],"performance":[{"prisoners":1}],"pete":[{"prisoners":1}],"phosphorescence":[{"prisoners":1}],"phosphorescent":[{"prisoners":1}],"physical":[{"prisoners":2}],"pick":[{"prisoners":2}],"picked":[{"prisoners":1}],"picture":[{"prisoners":2}],"pictured":[{"prisoners":1}],"pieces":[{"prisoners":1}],"pierced":[{"prisoners":1}],"pigeon":[{"prisoners":1}],"pike":[{"hard":1}],"pilot":[{"prisoners":1}],"pine":[{"hard":1}],"pines":[{"hard":1}],"pitch":[{"prisoners":1}],"pitched":[{"prisoners":2}],"pitiful":[{"prisoners":1}],"pitz":[{"hard":1}],"place":[{"prisoners":1}],"planet":[{"prisoners":4}],"planets":[{"prisoners":1}],"plans":[{"prisoners":1}],"plastic":[{"hard":1}],"pleased":[{"prisoners":1}],"plentiful":[{"prisoners":1}],"plenty":[{"prisoners":1}],"plodded":[{"prisoners":3}],"plunged":[{"prisoners":1}],"plunging":[{"hard":1}],"plutonium":[{"prisoners":1}],"pneumatic":[{"hard":1}],"pockets":[{"prisoners":1}],"point":[{"prisoners":1}],"polluted":[{"prisoners":2}],"porthole":[{"prisoners":1}],"position":[{"prisoners":1}],"positron":[{"prisoners":3}],"powerful":[{"prisoners":3}],"precious":[{"prisoners":1}],"predicament":[{"prisoners":1}],"preparing":[{"prisoners":1}],"presence":[{"prisoners":1}],"prickly":[{"prisoners":1}],"primitive":[{"prisoners":2}],"probability":[{"prisoners":1}],"problem":[{"prisoners":1}],"prologue":[{"hard":1}],"prominent":[{"prisoners":1}],"prospect":[{"prisoners":2}],"protruded":[{"prisoners":1}],"protruding":[{"prisoners":1}],"pyrophage":[{"prisoners":1}],"radiation":[{"prisoners":1}],"radioactive":[{"prisoners":5}],"radioactivity":[{"prisoners":1}],"rained":[{"prisoners":1}],"raising":[{"prisoners":1}],"random":[{"prisoners":2}],"rapidly":[{"prisoners":2}],"raw":[{"prisoners":1}],"reach":[{"prisoners":1}],"reaching":[{"prisoners":1}],"reactor":[{"prisoners":1}],"reactors":[{"prisoners":1}],"real":[{"prisoners":1}],"realized":[{"prisoners":5}],"reconnaissance":[{"prisoners":1}],"reconstruct":[{"prisoners":1}],"red":[{"prisoners":2}],"reflection":[{"hard":1}],"refraction":[{"prisoners":1}],"region":[{"prisoners":1}],"remains":[{"prisoners":1}],"repair":[{"prisoners":1}],"reply":[{"hard":1}],"reported":[{"hard":1}],"responsible":[{"prisoners":1}],"rested":[{"prisoners":2}],"retreated":[{"prisoners":1}],"return":[{"prisoners":1}],"returned":[{"prisoners":2}],"rifle":[{"hard":1}],"ring":[{"prisoners":1}],"rising":[{"prisoners":1}],"river":[{"prisoners":12}],"riverbank":[{"prisoners":1}],"riveted":[{"prisoners":1}],"road":[{"prisoners":12}],"roads":[{"prisoners":1}],"roaming":[{"prisoners":1}],"roaring":[{"prisoners":2}],"robinson":[{"prisoners":2}],"robots":[{"prisoners":1}],"roofs":[{"prisoners":1}],"roots":[{"hard":1}],"rose":[{"prisoners":2}],"routine":[{"prisoners":2}],"rudder":[{"hard":1}],"rumble":[{"prisoners":1}],"rumbling":[{"prisoners":3}],"running":[{"prisoners":1}],"rush":[{"hard":1}],"rushed":[{"prisoners":1}],"rusted":[{"prisoners":1}],"rustled":[{"prisoners":2}],"rusty":[{"prisoners":2}],"ruts":[{"prisoners":1}],"samples":[{"prisoners":1}],"sand":[{"prisoners":3}],"sandy":[{"hard":1}],"sang":[{"hard":1}],"sat":[{"hard":1},{"prisoners":1}],"savagely":[{"prisoners":1}],"scanned":[{"prisoners":1}],"scarcely":[{"prisoners":1}],"school":[{"prisoners":1}],"schools":[{"hard":1}],"scooping":[{"prisoners":1}],"scorching":[{"prisoners":2}],"scrambled":[{"prisoners":2}],"seadog":[{"hard":1}],"seconds":[{"prisoners":1}],"security":[{"prisoners":1}],"seeds":[{"prisoners":1}],"sender":[{"prisoners":2}],"senders":[{"prisoners":1}],"sensed":[{"prisoners":2}],"sensory":[{"prisoners":1}],"separated":[{"prisoners":1}],"serene":[{"prisoners":1}],"serious":[{"prisoners":1}],"shabby":[{"prisoners":1}],"shadow":[{"prisoners":1}],"shadows":[{"prisoners":1}],"shape":[{"prisoners":2}],"sharks":[{"hard":1}],"sharply":[{"prisoners":1}],"shell":[{"prisoners":1}],"shielded":[{"prisoners":1}],"ship":[{"prisoners":13}],"shook":[{"prisoners":1}],"shore":[{"hard":3},{"prisoners":1}],"shortening":[{"prisoners":1}],"shorts":[{"prisoners":2}],"shoulder":[{"prisoners":1}],"shoulders":[{"prisoners":1}],"showered":[{"prisoners":1}],"side":[{"hard":3},{"prisoners":3}],"sides":[{"prisoners":2}],"signal":[{"prisoners":1}],"silently":[{"hard":1}],"silhouetted":[{"prisoners":1}],"sinew":[{"hard":1}],"single":[{"hard":1}],"sizzling":[{"prisoners":1}],"sky":[{"hard":1},{"prisoners":4}],"slabs":[{"prisoners":2}],"sleeping":[{"prisoners":1}],"sliding":[{"hard":1}],"slope":[{"prisoners":1}],"sloped":[{"prisoners":1}],"sloping":[{"prisoners":1}],"slowed":[{"prisoners":1}],"sluggish":[{"prisoners":1}],"small":[{"prisoners":3}],"smelled":[{"prisoners":2}],"smoke":[{"prisoners":1}],"smoky":[{"prisoners":1}],"smooth":[{"prisoners":1}],"snack":[{"prisoners":1}],"sneakers":[{"prisoners":1}],"soft":[{"prisoners":1}],"solid":[{"prisoners":2}],"son":[{"prisoners":1}],"sooner":[{"prisoners":1}],"sooty":[{"prisoners":1}],"sorcerers":[{"prisoners":1}],"sort":[{"prisoners":2}],"sorts":[{"prisoners":1}],"space":[{"prisoners":3}],"spanned":[{"prisoners":1}],"sparse":[{"prisoners":2}],"special":[{"prisoners":1}],"sped":[{"prisoners":1}],"spell":[{"prisoners":1}],"spend":[{"prisoners":1}],"spiral":[{"prisoners":1}],"spot":[{"prisoners":1}],"spray":[{"prisoners":1}],"squat":[{"prisoners":1}],"squatted":[{"prisoners":2}],"squeamishly":[{"prisoners":1}],"stab":[{"prisoners":2}],"stage":[{"prisoners":1}],"standing":[{"prisoners":1}],"started":[{"prisoners":1}],"starve":[{"prisoners":1}],"stay":[{"hard":1}],"steady":[{"prisoners":1}],"steam":[{"prisoners":1}],"steel":[{"hard":1}],"steep":[{"prisoners":1}],"stepping":[{"prisoners":1}],"stifling":[{"prisoners":1}],"stinking":[{"prisoners":1}],"stock":[{"hard":1}],"stood":[{"hard":1}],"stopped":[{"prisoners":2}],"straight":[{"prisoners":1}],"strange":[{"prisoners":2}],"stratosphere":[{"prisoners":2}],"strewn":[{"prisoners":1}],"strings":[{"hard":1}],"struck":[{"prisoners":1}],"stuck":[{"prisoners":1}],"stuff":[{"prisoners":1}],"stuffed":[{"prisoners":1}],"stunned":[{"prisoners":1}],"stupid":[{"prisoners":1}],"style":[{"hard":1}],"subject":[{"prisoners":1}],"submerged":[{"prisoners":1}],"subsiding":[{"prisoners":1}],"substances":[{"prisoners":1}],"suddenly":[{"prisoners":4}],"suggestive":[{"prisoners":1}],"sulked":[{"prisoners":1}],"summon":[{"prisoners":1}],"sun":[{"hard":1}],"sunk":[{"prisoners":1}],"supposed":[{"prisoners":1}],"surely":[{"prisoners":1}],"surfaced":[{"prisoners":1}],"surprised":[{"prisoners":1}],"surrounding":[{"prisoners":1}],"survive":[{"prisoners":1}],"sustained":[{"prisoners":1}],"swarm":[{"prisoners":1}],"swayed":[{"hard":1}],"swim":[{"hard":1}],"swimming":[{"prisoners":1}],"swing":[{"prisoners":1}],"swirling":[{"prisoners":1}],"switched":[{"prisoners":1}],"swum":[{"prisoners":1}],"swung":[{"prisoners":1}],"system":[{"prisoners":1}],"t":[{"hard":2},{"prisoners":19}],"tail":[{"prisoners":1}],"tale":[{"prisoners":2}],"talents":[{"prisoners":1}],"tall":[{"prisoners":2}],"taller":[{"prisoners":1}],"tanks":[{"prisoners":1}],"tarkypark":[{"hard":1}],"teacher":[{"prisoners":1}],"technology":[{"hard":1}],"ten":[{"prisoners":1}],"terrain":[{"prisoners":1}],"thanked":[{"prisoners":1}],"theoretically":[{"prisoners":1}],"thick":[{"prisoners":1}],"thing":[{"prisoners":3}],"things":[{"prisoners":1}],"thinking":[{"prisoners":1}],"thirty":[{"prisoners":1}],"thought":[{"hard":1},{"prisoners":5}],"three":[{"prisoners":1}],"thudded":[{"prisoners":1}],"thunderous":[{"prisoners":1}],"time":[{"prisoners":5}],"times":[{"prisoners":2}],"timid":[{"prisoners":1}],"tire":[{"prisoners":1}],"toed":[{"prisoners":1}],"toes":[{"prisoners":1}],"told":[{"prisoners":3}],"tons":[{"prisoners":1}],"tossing":[{"prisoners":1}],"totz":[{"hard":1}],"trace":[{"prisoners":1}],"traces":[{"prisoners":1}],"tramping":[{"prisoners":1}],"trampled":[{"prisoners":1}],"transmitter":[{"prisoners":4}],"transparency":[{"prisoners":1}],"trap":[{"prisoners":1}],"treads":[{"prisoners":1}],"trees":[{"hard":1},{"prisoners":4}],"trusswork":[{"prisoners":1}],"trust":[{"hard":1}],"turn":[{"prisoners":3}],"turned":[{"prisoners":1}],"turning":[{"prisoners":1}],"twenty":[{"prisoners":4}],"twisted":[{"prisoners":1}],"types":[{"prisoners":1}],"ugly":[{"prisoners":4}],"uncivilized":[{"prisoners":1}],"uncomfortable":[{"prisoners":1}],"understood":[{"prisoners":1}],"undoubtedly":[{"prisoners":3}],"unexpected":[{"prisoners":1}],"unexplored":[{"prisoners":2}],"unit":[{"prisoners":1}],"unpleasant":[{"prisoners":2}],"unreal":[{"prisoners":1}],"uphill":[{"prisoners":1}],"upstream":[{"prisoners":1}],"upward":[{"prisoners":1}],"usual":[{"prisoners":1}],"vaguest":[{"prisoners":1}],"vanish":[{"prisoners":1}],"ve":[{"prisoners":3}],"vegetation":[{"prisoners":2}],"vicious":[{"prisoners":1}],"vines":[{"prisoners":1}],"violet":[{"prisoners":1}],"visible":[{"prisoners":2}],"visualize":[{"prisoners":1}],"vividly":[{"prisoners":1}],"voices":[{"prisoners":1}],"waist":[{"prisoners":1}],"waking":[{"prisoners":1}],"walked":[{"prisoners":3}],"walking":[{"prisoners":1}],"walls":[{"prisoners":1}],"washed":[{"prisoners":1}],"wasn":[{"prisoners":2}],"wastes":[{"prisoners":1}],"watched":[{"prisoners":1}],"water":[{"hard":2},{"prisoners":6}],"waterfall":[{"prisoners":1}],"waters":[{"prisoners":1}],"well":[{"prisoners":6}],"west":[{"prisoners":1}],"westward":[{"prisoners":1}],"wheel":[{"hard":1}],"wheezed":[{"prisoners":1}],"whine":[{"prisoners":1}],"whined":[{"prisoners":1}],"white":[{"prisoners":2}],"wide":[{"prisoners":1}],"wild":[{"prisoners":2}],"wilder":[{"prisoners":1}],"windows":[{"prisoners":1}],"wise":[{"prisoners":1}],"wiser":[{"prisoners":1}],"wolf":[{"prisoners":1}],"won":[{"prisoners":2}],"woods":[{"prisoners":3}],"work":[{"hard":1}],"worlds":[{"prisoners":2}],"wouldn":[{"prisoners":4}],"wound":[{"hard":1}],"wounded":[{"prisoners":1}],"wrapped":[{"prisoners":1}],"wriggling":[{"hard":1}],"year":[{"prisoners":1}],"years":[{"prisoners":2}],"yellow":[{"hard":3}],"yoke":[{"hard":1}],"yonder":[{"prisoners":1}],"young":[{"prisoners":1}],"zenith":[{"prisoners":2}]}`
	expected := Freq{"hard": 2, "prisoners": 22}

	idx := ReverseIdx{}
	idx.Init()

	err := json.Unmarshal([]byte(idxMap), &idx.M)
	if err != nil {
		t.Fatal(err)
	}

	actual := idx.Search("forest grass dark")
	if !reflect.DeepEqual(expected, actual) {
		t.Fail()
	}
}
