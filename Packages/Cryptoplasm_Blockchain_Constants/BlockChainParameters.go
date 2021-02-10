package Cryptoplasm_Blockchain_Constants

import (
    p "Cryptoplasm-Core/Packages/Firefly_Precision"
    "fmt"
    "time"
)
//
//	BlockChainParameters.go		                Cryptoplasm Blockchain Parameters variable and constants
//
//================================================
// Function List
//
//	01  - CryptoplasmDecimalSeedPrinter		prints the Decimal Seeds
//      02  - CryptoplasmIntegerSeedPrinter             prints the Integer Seeds
//      03  - CryptoplasmIntervals                      prints Cryptoplasm time intervals
//      03b - ConvertIntervals                          converts interval duration in time duration.
//
//=================================================================================================

//BlockChain Intervals
const (
    // Emission Details
    //
    // Founders-Amount (how much goes to Funding members) and InitialQ-Amount
    // (how much remains in the Conqueror Pool from the Premine after Founders-Amount is spent)
    // added together give the PreMine Amount
    //=============================================================================================
    TheoreticalEraEmission 	= "100000000000000.000000000000000000000000" //Total Theoretical Emission without Bonuses for One Era
    PreMinedAmount 		= 	"500000000.000000000000000000000000" //Premine Amount
    TheoreticalStartBlockReward = 	      "100.000000000000000000000000" //Base Block-Reward used as starting point for BR calculations
    AuPerUnit 			= 		 "1000000000000000000000000" // Atomic Units per Cryptoplasm
    //===============================
    //
    // Decimal Seed Precision
    //=============================================================================================
    CryptoplasmContextPrecision = 50
    CryptoplasmSeedPrecision    = uint32(40)
    CryptoplasmCurrencyPrecision= uint32(24)
    //===============================
    //
    // Fee per Byte
    //=============================================================================================
    MaxFeePerByte               = 		"0.000010000000000000000000" 			//10**19 AU aka 10 MicroPlasm aka 10000 NanoPlasm
    FeePerByteDecrement   	= 		"0.000000000000020000000000" 			//2*10**10 AU aka 20 FemtoPlasm per Block
    MinFeePerByte         	= 		"0.000000001000000000000000" 			//10**15 AU aka 1 NanoPlasm
    FeePerByteBHThreshold 	= 	"499950000"
    //===============================
    //
    // Transaction Fee Multipliers
    //=============================================================================================
    GAS    			= 1
    PLASMA 			= 3
    MIASMA 			= 5
    BLINK  			= 7
    OUTPUT 			= 400
    //===============================
    //
    // Interval Length
    //=============================================================================================
    IntervalPurple 		= "637"       	//Purple	Period in Blocks
    IntervalIndigo 		= "4459"      	//Indigo	Period in Blocks
    IntervalBlue   		= "31213"     	//Blue		Period in Blocks
    IntervalGreen  		= "218491"    	//Green		Period in Blocks
    IntervalYellow 		= "1529437"   	//Yellow	Period in Blocks
    IntervalOrange 		= "10706059"  	//Orange	Period in Blocks
    IntervalRed    		= "74942413"  	//Red		Period in Blocks
    IntervalWhite  		= "524596891" 	//White		One Era
    //===============================
    //
    // 1st,2nd and 3rd Seeds (Primary Decimal Seeds)
    //=============================================================================================
    BaseAreaSeed      		= 		"0.7412804842563701273935785745772150152399" 	//23115732923502492351869403393438306982639378690095569104032787203980148228041913685525272242174661218702645141327232481974409791932000078410388386177817655043149469072067"
    BaseHeightSeed    		= 		"1.4825609685127402547871571491544300304798" 	//46231465847004984703738806786876613965278757380191138208065574407960296456083827371050544484349322437405290282654464963948819583864000156820776772355635310086298938144135"
    CPAreaRatioSeed   		= 		"0.0061411618911384368190200232738753822284" 	//53558150274373282200996197268636748422270727410231633348787988192143721376975759958017147211196311637271799923302298183577206534981853131081653660870758170105689380545402"
    TotalComputedSum  		=  "99999499999999.999999999999999737703204"    	     	//at higher precision sum ends in ...205 ??
    HighestComputedBR 		= 	   "341281.048829773462213727839038"
    //===============================
    //
    // Cryptoplasm Decimal Seeds (Secondary Decimal Seeds)
    //=============================================================================================
    PurpleUpHeightSeed   	= 	      "909.7302802947269378763139749440449711467316" 	//26891978265178712859078440044665328695481061733347841808882751709087041846467952197400245227340660911511048676235584758898987802431939330594135241876828897943372661795441"
    PurpleUpAreaSeed     	= 	   "300586.0540637279168676302599502125746465902973" 	//75632644992903257458735950465103568423241907148623679450703441583915996934435313769203007730354943771681360454686785107877436679708139583522254486775190587479608684256174"
    PurpleDownHeightSeed 	=             "732.7846753997830870932228632878163206032142" 	//35660083738592050362142901663713847113432794566069457416419489439272636989036672482046638214546024574094739007826821411000258383259087452574481848085511568932620265261545"
    PurpleDownAreaSeed   	= 	   "293237.9599065836066898482542079790576762358547" 	//24432938779590995837841532019384790150349085420733522736133724177165438818776167178359672061917060017698829380688343277961621188229780984435117506605176498547028399816297"
    //===============================
    IndigoUpHeightPeakSeed   	= 	      "458.9111778370467284362043581893618334806483" 	//75673341438198016253867830693417334152383470765550007024833715596107482766988608639015947529510256007991003527898025942204481992736070083645428346888351518535457823580472"
    IndigoUpAreaPeakSeed     	= 	   "261044.7466542166158914952564909348348828575761" 	//47170870024639515423665252327046908149680993950708042561330034560537873507717108724954307264259363897304025369848969984617755648586455691736830848158213038771190145191942"
    IndigoDownHeightPeakSeed 	= 	      "381.8064188746302072671293276020255040603225" 	//77776097557209956537647420738431173891021295054572742092639406064277097443334557475597451949849357602982895378137618856526279561729612727096614403090545562180108554824788"
    IndigoDownAreaPeakSeed   	= 	   "247417.8525721708221416643020274773849304925499" 	//77916987540693469736206911373031390750240324496266620224972229908647554039298382508242888156344262818782463224905235604572942254400321050377665099311491108148214911008294"
    //===============================
    BlueUpHeightPeakSeed   	= 	      "478.9040423687162704145585733430684419287750" 	//32592232455024302156582579310007666023578836271607375165152790420196164894495691344344356098544713341783994549774764791511458586763746141409119821970206957891019496013170"
    BlueUpAreaPeakSeed     	= 	   "264250.7806520145541428342346114436078361575829" 	//85896656772099968049965664720775391354733769891522665841287865279381182447309216963759770527490938679403384007259874973254896360573674241631789715365733742006539045082405"
    BlueDownHeightPeakSeed 	= 	       "60.4801406253969681955672975171257116900208" 	//35762604185541729437790771581361892297073237879616129659398556012213354164373323985040401887511553112675019058432069804561867343974960641233072376535458128585610652668701"
    BlueDownAreaPeakSeed   	= 	   "169040.4328894502748717817530378647839451601223" 	//14469301683209967785660336971307130171983630485525806795551381667712795800898521584265423182588320749466935042998433502899035050300320599018761484513376903429304753104722"
    //===============================
    GreenUpHeightPeakSeed   	= 	      "111.8438950170378749137571471887605954622164" 	//12867757092623184799536723723984817254131297365666777161093219085498324730212978578142100085160445949651383195676180265136049669100834118651320683764852953446043757242434"
    GreenUpAreaPeakSeed    	= 	   "183907.2346248548867458858193691529488817424623" 	//39113027015993818730764424719296504565018872998656781769837509572549708407470164860932029698360333725389480068793454194898480993196326213690435887363783841023900191977303"
    GreenDownHeightPeakSeed 	= 	      "432.0518773260566120066094740814362455715120" 	//03866539218214160281019712693375711308362670711363704454595128840820882089215316016167870115360954967712081058848109623142262259194493083293911074217930769540096024450561"
    GreenDownAreaPeakSeed   	= 	   "256525.4153140180257492988414765980164505017216" 	//17603469406324699264552203728228781734854047563080266864889917379915522679125848517743518074977100995657589506504988514221772587652078403535578174946110910877306175175831"
    //===============================
    YellowUpHeightPeakSeed   	= 	      "289.7729212414708818791070741090082183417430" 	//05493868213596096365590475181833258358465672817908846494426160474614574654369161757532776621391535689893000433337201721222100282080398865859651228532805344145078971867156"
    YellowUpAreaPeakSeed     	= 	   "228527.6547124002743539768227771865451661334612" 	//82051023429631295084228214141774093090232076185162203767778977868683417696291068680659145751642959282977468816748030564224610991180351121691818457543376840110408220056801"
    YellowDownHeightPeakSeed 	= 	      "269.3351469003894115627079132197617241905450" 	//18396904506226212361196557891241987452343475676335561525682588401208319343318545668692751385857807667204710244635373112015235949039067998472592935687779763040136503026148"
    YellowDownAreaPeakSeed   	= 	   "223945.1130637908409935147320368856678977192584" 	//36541018240740246649576735536180292375694773812166491930655995247799212452046610949323116974900711193004663009289356729830960286123820285321368328210923027750793180868603"
    //===============================
    OrangeUpHeightPeakSeed   	=             "179.7509619098377040536257598287503155662095" 	//72789615599211546246855331394309488651847507972496123667662554571470524011673280927601614309076246446328521090638757140548152543738622480383637737592709642333820953400965"
    OrangeUpAreaPeakSeed     	= 	   "202196.5903798765616686986828041481174745740673" 	//03702616454709778156777495028153971543184141005677899889589499007422130010346002126582056241278715116914223405977394279769720179755860307152238320388967160138703761983160"
    OrangeDownHeightPeakSeed 	= 	      "280.4243040081237796219338046539550049529617" 	//55541775726452733784575987793582318465163411038588984055513086710270571398803499810759402046824905309077200291776303022464900114246823372595531758644978016364890570530864"
    OrangeDownAreaPeakSeed   	= 	   "226448.9964125797532724906559614943733250122097" 	//81011964613644048747402037909317201310192096624981879288619518040599998635956022076851839404542253563216844707158437345374885758966190191892475017804054990945748522013174"
    //===============================
    RedUpHeightPeakSeed   	= 	      "185.2707213895297693451326061147780910517883" 	//71580515908879151284881604632153356843246622823244290161927325938691182581452507122203432887230124264762661404338862761784959484426833352783146632351792561108840033909110"
    RedUpAreaPeakSeed     	= 	   "203614.8784555785585271967245268050063415061685" 	//77613928439143777266907000301408484507837129473340838769099596496184639293804298777173098509834768715829137756304537771292676634707757587159664600859089495362064618633161"
    RedDownHeightPeakSeed 	= 	      "561.2112652145994814286298512397167158889632" 	//82040880690766415114927327050916082448016618724791510731215394255084134213230824510555923992158583594452586869208051730612161579190252612660288257088154809059958316016090"
    RedDownAreaPeakSeed   	= 	   "276029.6308180069483705972098155793017931637200" 	//12537782878339614625771498809035880098867872346732897743180678569403622937685906472124480697705107611139512091414204363056523843258830147818880718433631053114090520642197"
    //===============================
    WhiteHeightPeakSeed 	= 	      "380.1281559433085894190653590677886632087970" 	//64030092438005383734298957536370484716894250611339362185936095502532022628114548971046078379731970307243636309202148825978990235984310581518536951072236953890507709658038"
    WhiteAreaPeakSeed   	= 	   "247098.9538162242103958865943966608833446133772" 	//96509621441960570900115269772036588799224478842644813092952995468374792700535540513090096414588656427333291842393759390325440824474269113899493647236888735135953137095747"
    //===============================
    //
    // Cryptoplasm Integer Seeds (Integer Seeds)
    //=============================================================================================
    PurpleUpPeakSeed   		= 626
    PurpleDownPeakSeed 		= 72
    IndigoUpPeakSeed   		= 474
    IndigoDownPeakSeed 		= 448
    BlueUpPeakSeed     		= 157
    BlueDownPeakSeed   		= 339
    GreenUpPeakSeed    		= 357
    GreenDownPeakSeed  		= 465
    YellowUpPeakSeed   		= 417
    YellowDownPeakSeed 		= 410
    OrangeUpPeakSeed   		= 380
    OrangeDownPeakSeed 		= 414
    RedUpPeakSeed      		= 381
    RedDownPeakSeed    		= 508
    WhitePeakSeed      		= 191
)

var (
    //
    // Cryptoplasm p Decimal Context Definition
    //=============================================================================================
    CryptoplasmPrecisionContext 	= p.Context{
        Precision:   		CryptoplasmContextPrecision,
        MaxExponent: 		p.MaxExponent,
        MinExponent: 		p.MinExponent,
        Rounding:    		p.RoundDown,
        // Default error conditions.
        Traps: 			p.InvalidOperation,
    }
    //
    //
    // Blockchain Premine and Emission variable Declarations
    //=============================================================================================
    TotalEmission 		= p.NFS(TheoreticalEraEmission)
    PreMine       		= p.NFS(PreMinedAmount)
    StartBRd     		= p.NFS(TheoreticalStartBlockReward)
    AUs          		= p.NFS(AuPerUnit)
    CamelSupportReward  	= MULcp(White, StartBRd)                            		//Sum of Reward equivalent to the Camel-Support-Area
    CamelReward         	= DIFcp(TotalEmission, PreMine, CamelSupportReward) 		//Sum of Reward equivalent to the Camel Emission only (without its Support)
    //===============================
    //
    // Fee per byte variable Declaration
    //
    FpBMax 			= p.NFS(MaxFeePerByte)
    FpBInt 			= p.NFS(FeePerByteDecrement)
    FpBMin 			= p.NFS(MinFeePerByte)
    FpBThr 			= p.NFS(FeePerByteBHThreshold)
    //===============================
    //
    // Intervals and Seeds variable Declaration
    //=============================================================================================
    Purple 			= p.NFS(IntervalPurple)
    Indigo 			= p.NFS(IntervalIndigo)
    Blue   			= p.NFS(IntervalBlue)
    Green  			= p.NFS(IntervalGreen)
    Yellow 			= p.NFS(IntervalYellow)
    Orange 			= p.NFS(IntervalOrange)
    Red   			= p.NFS(IntervalRed)
    White  			= p.NFS(IntervalWhite)
    //===============================
    Seed1st 		= p.NFS(BaseAreaSeed)
    Seed2nd 		= p.NFS(BaseHeightSeed)
    Seed3rd 		= p.NFS(CPAreaRatioSeed)
    Seed4th 		= p.NFS(TotalComputedSum)
    Seed5th 		= p.NFS(HighestComputedBR)
    //===============================
    Seed01 			= p.NFS(PurpleUpHeightSeed)
    Seed02 			= p.NFS(PurpleUpAreaSeed)
    Seed03 			= p.NFS(PurpleDownHeightSeed)
    Seed04 			= p.NFS(PurpleDownAreaSeed)
    //===============================
    Seed05 			= p.NFS(IndigoUpHeightPeakSeed)
    Seed06 			= p.NFS(IndigoUpAreaPeakSeed)
    Seed07 			= p.NFS(IndigoDownHeightPeakSeed)
    Seed08 			= p.NFS(IndigoDownAreaPeakSeed)
    //===============================
    Seed09 			= p.NFS(BlueUpHeightPeakSeed)
    Seed10 			= p.NFS(BlueUpAreaPeakSeed)
    Seed11 			= p.NFS(BlueDownHeightPeakSeed)
    Seed12 			= p.NFS(BlueDownAreaPeakSeed)
    //===============================
    Seed13 			= p.NFS(GreenUpHeightPeakSeed)
    Seed14 			= p.NFS(GreenUpAreaPeakSeed)
    Seed15 			= p.NFS(GreenDownHeightPeakSeed)
    Seed16 			= p.NFS(GreenDownAreaPeakSeed)
    //===============================
    Seed17 			= p.NFS(YellowUpHeightPeakSeed)
    Seed18 			= p.NFS(YellowUpAreaPeakSeed)
    Seed19 			= p.NFS(YellowDownHeightPeakSeed)
    Seed20 			= p.NFS(YellowDownAreaPeakSeed)
    //===============================
    Seed21 			= p.NFS(OrangeUpHeightPeakSeed)
    Seed22 			= p.NFS(OrangeUpAreaPeakSeed)
    Seed23 			= p.NFS(OrangeDownHeightPeakSeed)
    Seed24 			= p.NFS(OrangeDownAreaPeakSeed)
    //===============================
    Seed25 			= p.NFS(RedUpHeightPeakSeed)
    Seed26 			= p.NFS(RedUpAreaPeakSeed)
    Seed27 			= p.NFS(RedDownHeightPeakSeed)
    Seed28 			= p.NFS(RedDownAreaPeakSeed)
    //===============================
    Seed29 			= p.NFS(WhiteHeightPeakSeed)
    Seed30 			= p.NFS(WhiteAreaPeakSeed)
    //===============================
    Seed31 			= PurpleUpPeakSeed
    Seed32 			= PurpleDownPeakSeed
    Seed33 			= IndigoUpPeakSeed
    Seed34 			= IndigoDownPeakSeed
    Seed35 			= BlueUpPeakSeed
    Seed36 			= BlueDownPeakSeed
    Seed37 			= GreenUpPeakSeed
    Seed38 			= GreenDownPeakSeed
    Seed39 			= YellowUpPeakSeed
    Seed40 			= YellowDownPeakSeed
    Seed41 			= OrangeUpPeakSeed
    Seed42 			= OrangeDownPeakSeed
    Seed43 			= RedUpPeakSeed
    Seed44 			= RedDownPeakSeed
    Seed45 			= WhitePeakSeed
    //===============================
    //
    // Cryptoplasm Geometric Height Intervals
    //
    //=============================================================================================
    ngSeed03 			= SUBcp(p.NFI(0), Seed03)
    ngSeed09 			= SUBcp(p.NFI(0), Seed09)
    ngSeed29 			= SUBcp(p.NFI(0), Seed29)
    //
    // Geometric Heights List Definition
    //=============================================================================================
    Ax 				= CryptoplasmSecondaryGeometricListing(Seed31, Seed01, Seed02)
    Bx 				= CryptoplasmSecondaryGeometricListing(Seed32, ngSeed03, Seed04)
    Cx 				= CryptoplasmSecondaryGeometricListing(Seed33, Seed05, Seed06)
    Dx 				= CryptoplasmSecondaryGeometricListing(Seed34, Seed07, Seed08)
    Ex 				= CryptoplasmSecondaryGeometricListing(Seed35, ngSeed09, Seed10)
    Fx 				= CryptoplasmSecondaryGeometricListing(Seed36, Seed11, Seed12)
    Gx 				= CryptoplasmSecondaryGeometricListing(Seed37, Seed13, Seed14)
    Hx 				= CryptoplasmSecondaryGeometricListing(Seed38, Seed15, Seed16)
    Ix 				= CryptoplasmSecondaryGeometricListing(Seed39, Seed17, Seed18)
    Jx 				= CryptoplasmSecondaryGeometricListing(Seed40, Seed19, Seed20)
    Kx 				= CryptoplasmSecondaryGeometricListing(Seed41, Seed21, Seed22)
    Lx 				= CryptoplasmSecondaryGeometricListing(Seed42, Seed23, Seed24)
    Mx 				= CryptoplasmSecondaryGeometricListing(Seed43, Seed25, Seed26)
    Nx 				= CryptoplasmSecondaryGeometricListing(Seed44, Seed27, Seed28)
    Ox 				= CryptoplasmSecondaryGeometricListing(Seed45, ngSeed29, Seed30)
    //===============================
    //
    // Computing Fee Matrix
    //=============================================================================================
    FE01 			= GAS
    FE02 			= GAS * BLINK
    FE03 			= GAS * OUTPUT
    FE04 			= GAS * OUTPUT * BLINK
    //===============================
    FE05 			= FE01 * PLASMA
    FE06 			= FE02 * PLASMA
    FE07 			= FE03 * PLASMA
    FE08 			= FE04 * PLASMA
    //===============================
    FE09 			= FE01 * MIASMA
    FE10 			= FE02 * MIASMA
    FE11 			= FE03 * MIASMA
    FE12 			= FE04 * MIASMA
    //===============================
    ArrayGAS 			= [...]int{FE01, FE02, FE03, FE04} 		//Array declarations
    ArrayPLS 			= [...]int{FE05, FE06, FE07, FE08} 		//Array declarations
    ArrayMSM 			= [...]int{FE09, FE10, FE11, FE12} 		//Array declarations
    //===============================
    FeeArray 			= [3][4]int{ArrayGAS, ArrayPLS, ArrayMSM} 	//Array of Array declaration
    //===============================
    //This is how the Fee Array should look like:
    //Gasses_List = [ 1,  7,  400, 2800]
    //Plasma_List = [ 3, 21, 1200, 8400]
    //Miasma_List = [ 5, 35, 2000,14000]

)
//================================================
//
// Function 01 - CryptoplasmDecimalSeedPrinter
//
// CryptoplasmDecimalSeedPrinter prints all the Decimal Seeds
func CryptoplasmDecimalSeedPrinter() {
    start := time.Now()
    fmt.Println("=======PRIMARY DECIMAL SEEDS============================")
    fmt.Println("Seed1st, ", BaseAreaSeed)
    fmt.Println("Seed2nd, ", BaseHeightSeed)
    fmt.Println("Seed3nd, ", CPAreaRatioSeed)
    fmt.Println("=======SECONDARY DECIMAL SEEDS==========================")
    fmt.Println("PurpleUpHeightSeed is:         ", PurpleUpHeightSeed)
    fmt.Println("PurpleUpAreaSeed is:        ", PurpleUpAreaSeed)
    fmt.Println("PurpleDownHeightSeed is:       ", PurpleDownHeightSeed)
    fmt.Println("PurpleDownAreaSeed is:      ", PurpleDownAreaSeed)
    fmt.Println("========================================================")
    fmt.Println("IndigoUpHeightPeakSeed is:     ", IndigoUpHeightPeakSeed)
    fmt.Println("IndigoUpAreaPeakSeed is:    ", IndigoUpAreaPeakSeed)
    fmt.Println("IndigoDownHeightPeakSeed is:   ", IndigoDownHeightPeakSeed)
    fmt.Println("IndigoDownAreaPeakSeed is:  ", IndigoDownAreaPeakSeed)
    fmt.Println("========================================================")
    fmt.Println("BlueUpHeightPeakSeed is:       ", BlueUpHeightPeakSeed)
    fmt.Println("BlueUpAreaPeakSeed is:      ", BlueUpAreaPeakSeed)
    fmt.Println("BlueDownHeightPeakSeed is:      ", BlueDownHeightPeakSeed)
    fmt.Println("BlueDownAreaPeakSeed is:    ", BlueDownAreaPeakSeed)
    fmt.Println("========================================================")
    fmt.Println("GreenUpHeightPeakSeed is:      ", GreenUpHeightPeakSeed)
    fmt.Println("GreenUpAreaPeakSeed is:     ", GreenUpAreaPeakSeed)
    fmt.Println("GreenDownHeightPeakSeed is:    ", GreenDownHeightPeakSeed)
    fmt.Println("GreenDownAreaPeakSeed is:   ", GreenDownAreaPeakSeed)
    fmt.Println("========================================================")
    fmt.Println("YellowUpHeightPeakSeed is:     ", YellowUpHeightPeakSeed)
    fmt.Println("YellowUpAreaPeakSeed is:    ", YellowUpAreaPeakSeed)
    fmt.Println("YellowDownHeightPeakSeed is:   ", YellowDownHeightPeakSeed)
    fmt.Println("YellowDownAreaPeakSeed is:  ", YellowDownAreaPeakSeed)
    fmt.Println("========================================================")
    fmt.Println("OrangeUpHeightPeakSeed is:     ", OrangeUpHeightPeakSeed)
    fmt.Println("OrangeUpAreaPeakSeed is:    ", OrangeUpAreaPeakSeed)
    fmt.Println("OrangeDownHeightPeakSeed:      ", OrangeDownHeightPeakSeed)
    fmt.Println("OrangeDownAreaPeakSeed is:  ", OrangeDownAreaPeakSeed)
    fmt.Println("========================================================")
    fmt.Println("RedUpHeightPeakSeed is:        ", RedUpHeightPeakSeed)
    fmt.Println("RedUpAreaPeakSeed is:       ", RedUpAreaPeakSeed)
    fmt.Println("RedDownHeightPeakSeed:         ", RedDownHeightPeakSeed)
    fmt.Println("RedDownAreaPeakSeed is:     ", RedDownAreaPeakSeed)
    fmt.Println("========================================================")
    fmt.Println("WhiteHeightPeakSeed is:        ", WhiteHeightPeakSeed)
    fmt.Println("WhiteAreaPeakSeed is:       ", WhiteAreaPeakSeed)
    fmt.Println("========================================================")
    elapsed := time.Since(start)
    fmt.Println("Printing took", elapsed)
}
//================================================
//
// Function 02 - CryptoplasmIntegerSeedPrinter
//
// CryptoplasmIntegerSeedPrinter prints all the Decimal Seeds
func CryptoplasmIntegerSeedPrinter() {
    start := time.Now()
    fmt.Println("=======INTEGER SEEDS====================================")
    fmt.Println("PurpleUpPeakSeed is:  ", PurpleUpPeakSeed)
    fmt.Println("PurpleDownPeakSeed is:", PurpleDownPeakSeed)
    fmt.Println("IndigoUpPeakSeed is:  ", IndigoUpPeakSeed)
    fmt.Println("IndigoDownPeakSeed is:", IndigoDownPeakSeed)
    fmt.Println("BlueUpPeakSeed is:    ", BlueUpPeakSeed)
    fmt.Println("BlueDownPeakSeed is:  ", BlueDownPeakSeed)
    fmt.Println("GreenUpPeakSeed is:   ", GreenUpPeakSeed)
    fmt.Println("GreenDownPeakSeed is: ", GreenDownPeakSeed)
    fmt.Println("YellowUpPeakSeed is:  ", YellowUpPeakSeed)
    fmt.Println("YellowDownPeakSeed is:", YellowDownPeakSeed)
    fmt.Println("OrangeUpPeakSeed is:  ", OrangeUpPeakSeed)
    fmt.Println("OrangeDownPeakSeed is:", OrangeDownPeakSeed)
    fmt.Println("RedUpPeakSeed is:     ", RedUpPeakSeed)
    fmt.Println("RedDownPeakSeed is:   ", RedDownPeakSeed)
    fmt.Println("WhitePeakSeed is:     ", WhitePeakSeed)
    fmt.Println("========================================================")
    elapsed := time.Since(start)
    fmt.Println("Printing took", elapsed)
}
//================================================
//
// Function 03 - CryptoplasmIntervals
//
// CryptoplasmIntervals prints the Cryptoplasm Time Intervals
func CryptoplasmIntervals() {
    start := time.Now()
    fmt.Println("=======Cryptoplasm Intervals======================================================")
    fmt.Println("Considering 1 Minute Block Time:")
    fmt.Println("")
    Y, D, H, m := ConvertIntervals(Purple)
    fmt.Println("Purple Interval is       ", Purple, "Blocks Long")
    fmt.Println("Purple Interval duration is:                ", Y, "years,  ", D, "days,", H, "hours,", m, "min.")
    fmt.Println("==================================================================================")
    Y, D, H, m = ConvertIntervals(Indigo)
    fmt.Println("Indigo Interval is      ", Indigo, "Blocks Long")
    fmt.Println("Indigo Interval duration is:                ", Y, "years,  ", D, "days, ", H, "hours,", m, "min.")
    fmt.Println("==================================================================================")
    Y, D, H, m = ConvertIntervals(Blue)
    fmt.Println("Blue Interval is       ", Blue, "Blocks Long")
    fmt.Println("Blue Interval duration is:                  ", Y, "years, ", D, "days,", H, "hours,", m, "min.")
    fmt.Println("==================================================================================")
    Y, D, H, m = ConvertIntervals(Green)
    fmt.Println("Green Interval is     ", Green, "Blocks Long")
    fmt.Println("Green Interval duration is:                 ", Y, "years,", D, "days,", H, "hours,", m, "min.")
    fmt.Println("==================================================================================")
    Y, D, H, m = ConvertIntervals(Yellow)
    fmt.Println("Yellow Interval is   ", Yellow, "Blocks Long")
    fmt.Println("Yellow Interval duration is:                ", Y, "years,", D, "days, ", H, "hours,", m, "min.")
    fmt.Println("==================================================================================")
    Y, D, H, m = ConvertIntervals(Orange)
    fmt.Println("Orange Interval is  ", Orange, "Blocks Long")
    fmt.Println("Orange Interval duration is:               ", Y, "years,", D, "days,", H, "hours,", m, "min.")
    fmt.Println("==================================================================================")
    Y, D, H, m = ConvertIntervals(Red)
    fmt.Println("Red Interval is     ", Red, "Blocks Long")
    fmt.Println("Red Interval duration is:                 ", Y, "years,", D, "days, ", H, "hours,", m, "min.")
    fmt.Println("==================================================================================")
    Y, D, H, m = ConvertIntervals(White)
    fmt.Println("White Interval is  ", White, "Blocks Long")
    fmt.Println("White Interval duration is:               ", Y, "years, ", D, "days, ", H, "hours,", m, "min.")
    fmt.Println("==================================================================================")
    elapsed := time.Since(start)
    fmt.Println("Printing took", elapsed)
}
//================================================
//
// Function 03b - ConvertIntervals
//
// ConvertIntervals converts interval duration in time duration
func ConvertIntervals(interval *p.Decimal) (*p.Decimal, *p.Decimal, *p.Decimal, *p.Decimal) {
    var (
        pHours = new(p.Decimal)
        pDays  = new(p.Decimal)
        pYears = new(p.Decimal)
    )

    Years := DIVcp(interval, p.NFI(525600))
    _, _ = c.Quantize(pYears, Years, 0)
    Rest := SUBcp(interval, MULcp(pYears, p.NFI(525600)))

    Days := DIVcp(Rest, p.NFI(1440))
    _, _ = c.Quantize(pDays, Days, 0)
    Rest = SUBcp(Rest, MULcp(pDays, p.NFI(1440)))

    Hours := DIVcp(Rest, p.NFI(60))
    _, _ = c.Quantize(pHours, Hours, 0)
    Rest = SUBcp(Rest, MULcp(pHours, p.NFI(60)))

    Minutes := Rest

    return pYears, pDays, pHours, Minutes
}
