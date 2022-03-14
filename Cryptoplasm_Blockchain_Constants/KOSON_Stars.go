package Cryptoplasm_Blockchain_Constants

import p "github.com/Crypt0plasm/Firefly-APD"

type StarStructure struct {
    Primary *p.Decimal
    Secondary *p.Decimal
    Tertiary *p.Decimal
}

var (
    GoldenStarDust 	= StarStructure{p.NFI(100),p.NFI(75),p.NFI(25)}
    SilverStarDust 	= StarStructure{p.NFI(20),p.NFI(15),p.NFI(5)}
    BronzeStarDust 	= StarStructure{p.NFI(4),p.NFI(3),p.NFI(1)}

    GoldenStarLight 	= StarStructure{p.NFI(196),p.NFI(147),p.NFI(49)}
    SilverStarLight 	= StarStructure{p.NFI(28),p.NFI(21),p.NFI(7)}
    BronzeStarLight 	= BronzeStarDust

    GoldenStarLightA 	= StarStructure{p.NFS("7.8400"),p.NFS("5.8800"),p.NFS("1.9600")}
    SilverStarLightA 	= StarStructure{p.NFS("0.2800"),p.NFS("0.2100"),p.NFS("0.0700")}
    BronzeStarLightA 	= StarStructure{p.NFS("0.0100"),p.NFS("0.0075"),p.NFS("0.0025")}
)

func StarDustGoldenComputer(Star StarStructure) (StarDust *p.Decimal) {
    	S1 := MULxc(Star.Primary,GoldenStarDust.Primary)
    	S2 := MULxc(Star.Secondary,GoldenStarDust.Secondary)
    	S3 := MULxc(Star.Tertiary,GoldenStarDust.Tertiary)
    	StarDust = SUMxc(S1,S2,S3)
	return StarDust
}

func StarDustSilverComputer(Star StarStructure) (StarDust *p.Decimal) {
	S1 := MULxc(Star.Primary,SilverStarDust.Primary)
	S2 := MULxc(Star.Secondary,SilverStarDust.Secondary)
	S3 := MULxc(Star.Tertiary,SilverStarDust.Tertiary)
	StarDust = SUMxc(S1,S2,S3)
	return StarDust
}

func StarDustBronzeComputer(Star StarStructure) (StarDust *p.Decimal) {
	S1 := MULxc(Star.Primary,BronzeStarDust.Primary)
	S2 := MULxc(Star.Secondary,BronzeStarDust.Secondary)
	S3 := MULxc(Star.Tertiary,BronzeStarDust.Tertiary)
	StarDust = SUMxc(S1,S2,S3)
	return StarDust
}

func StarDustComputer (Gold, Silver, Bronze StarStructure) (StarDust *p.Decimal) {
	GoldenDust := StarDustGoldenComputer(Gold)
	SilverDust := StarDustSilverComputer(Silver)
	BronzeDust := StarDustBronzeComputer(Bronze)
	StarDust = SUMxc(GoldenDust,SilverDust,BronzeDust)
	return StarDust
}

func StarLightComputer (Gold, Silver, Bronze StarStructure) (BaseStarLight, StarLightAmplificationFactor, StarLight *p.Decimal) {
    	G1 := MULxc(Gold.Primary,GoldenStarLight.Primary)				//GoldenInternalBaseSL
    	G2 := MULxc(SUBxc(Gold.Primary,p.NFI(1)),GoldenStarLightA.Primary)	//GoldenInternalIntrinsicSLAmplification
    	G3 := MULxc(Gold.Primary,GoldenStarLightA.Primary)				//GoldenInternalExtrinsicSLAmplification
    	G4 := MULxc(Gold.Secondary,GoldenStarLight.Secondary)				//GoldenExternalBaseSL
    	G5 := MULxc(SUBxc(Gold.Secondary,p.NFI(1)),GoldenStarLightA.Secondary)	//GoldenExternalIntrinsicSLAmplification
    	G6 := MULxc(Gold.Secondary,GoldenStarLightA.Secondary)				//GoldenExternalExtrinsicSLAmplification
    	G7 := MULxc(Gold.Tertiary,GoldenStarLight.Tertiary)				//GoldenRemainingBaseSL
    	G8 := MULxc(SUBxc(Gold.Tertiary,p.NFI(1)),GoldenStarLightA.Tertiary)	//GoldenRemainingIntrinsicSLAmplification
    	G9 := MULxc(Gold.Tertiary,GoldenStarLightA.Tertiary)				//GoldenRemainingExtrinsicSLAmplification


	S1 := MULxc(Silver.Primary,SilverStarLight.Primary)				//SilverInternalBaseSL
	S2 := MULxc(SUBxc(Silver.Primary,p.NFI(1)),SilverStarLightA.Primary)	//SilverInternalIntrinsicSLAmplification
	S3 := MULxc(Silver.Primary,SilverStarLightA.Primary)				//SilverInternalExtrinsicSLAmplification
	S4 := MULxc(Silver.Secondary,SilverStarLight.Secondary)				//SilverExternalBaseSL
	S5 := MULxc(SUBxc(Silver.Secondary,p.NFI(1)),SilverStarLightA.Secondary)	//SilverExternalIntrinsicSLAmplification
	S6 := MULxc(Silver.Secondary,SilverStarLightA.Secondary)			//SilverExternalExtrinsicSLAmplification
	S7 := MULxc(Silver.Tertiary,SilverStarLight.Tertiary)				//SilverRemainingBaseSL
	S8 := MULxc(SUBxc(Silver.Tertiary,p.NFI(1)),SilverStarLightA.Tertiary)	//SilverRemainingIntrinsicSLAmplification
	S9 := MULxc(Silver.Tertiary,SilverStarLightA.Tertiary)				//SilverRemainingExtrinsicSLAmplification


	B1 := MULxc(Bronze.Primary,BronzeStarLight.Primary)				//BronzeInternalBaseSL
	B2 := MULxc(SUBxc(Bronze.Primary,p.NFI(1)),BronzeStarLightA.Primary)	//BronzeInternalIntrinsicSLAmplification
	B3 := MULxc(Bronze.Primary,BronzeStarLightA.Primary)				//BronzeInternalExtrinsicSLAmplification
	B4 := MULxc(Bronze.Secondary,BronzeStarLight.Secondary)				//BronzeExternalBaseSL
	B5 := MULxc(SUBxc(Bronze.Secondary,p.NFI(1)),BronzeStarLightA.Secondary)	//BronzeExternalIntrinsicSLAmplification
	B6 := MULxc(Bronze.Secondary,BronzeStarLightA.Secondary)			//BronzeExternalExtrinsicSLAmplification
	B7 := MULxc(Bronze.Tertiary,BronzeStarLight.Tertiary)				//BronzeRemainingBaseSL
	B8 := MULxc(SUBxc(Bronze.Tertiary,p.NFI(1)),BronzeStarLightA.Tertiary)	//BronzeRemainingIntrinsicSLAmplification
	B9 := MULxc(Bronze.Tertiary,BronzeStarLightA.Tertiary)				//BronzeRemainingExtrinsicSLAmplification

	G10 := ADDxc(p.NFI(1),DIVxc(SUMxc(G2,G6,G9,S3,S6,S9,B3,B6,B9),p.NFI(100)))	//GoldenInternalTotalAmplification
    	G11 := ADDxc(p.NFI(1),DIVxc(SUMxc(G3,G5,G9,S3,S6,S9,B3,B6,B9),p.NFI(100)))	//GoldenExternalTotalAmplification
    	G12 := ADDxc(p.NFI(1),DIVxc(SUMxc(G3,G6,G8,S3,S6,S9,B3,B6,B9),p.NFI(100)))	//GoldenRemainingTotalAmplification
	S10 := ADDxc(p.NFI(1),DIVxc(SUMxc(G3,G6,G9,S2,S6,S9,B3,B6,B9),p.NFI(100)))	//SilverInternalTotalAmplification
	S11 := ADDxc(p.NFI(1),DIVxc(SUMxc(G3,G6,G9,S3,S5,S9,B3,B6,B9),p.NFI(100)))	//SilverExternalTotalAmplification
	S12 := ADDxc(p.NFI(1),DIVxc(SUMxc(G3,G6,G9,S3,S6,S8,B3,B6,B9),p.NFI(100)))	//SilverRemainingTotalAmplification
	B10 := ADDxc(p.NFI(1),DIVxc(SUMxc(G3,G6,G9,S3,S6,S9,B2,B6,B9),p.NFI(100)))	//BronzeInternalTotalAmplification
	B11 := ADDxc(p.NFI(1),DIVxc(SUMxc(G3,G6,G9,S3,S6,S9,B3,B5,B9),p.NFI(100)))	//BronzeExternalTotalAmplification
	B12 := ADDxc(p.NFI(1),DIVxc(SUMxc(G3,G6,G9,S3,S6,S9,B3,B6,B8),p.NFI(100)))	//BronzeRemainingTotalAmplification


	G13 := MULxc(G1,G10)								//GoldenInternalAmplifiedSL
    	G14 := MULxc(G4,G11)								//GoldenInternalAmplifiedSL
    	G15 := MULxc(G7,G12)								//GoldenInternalAmplifiedSL

	S13 := MULxc(S1,S10)								//SilverInternalAmplifiedSL
	S14 := MULxc(S4,S11)								//SilverInternalAmplifiedSL
	S15 := MULxc(S7,S12)								//SilverInternalAmplifiedSL

	B13 := MULxc(B1,B10)								//BronzeInternalAmplifiedSL
	B14 := MULxc(B4,B11)								//BronzeInternalAmplifiedSL
	B15 := MULxc(B7,B12)								//BronzeInternalAmplifiedSL

	BaseStarLight = SUMxc(G1,G4,G7,S1,S4,S7,B1,B4,B7)
	AmplifiedStarLight := SUMxc(G13,G14,G15,S13,S14,S15,B13,B14,B15)
	StarLightAmplificationFactor = DIVxc(AmplifiedStarLight,BaseStarLight)

	StarLight = DIVxc(AmplifiedStarLight,p.NFI(100))

	return BaseStarLight, StarLightAmplificationFactor, StarLight
}