package KosonicWritings

type Kanon struct {
    Name string
    Value string
}

var (
    KossonTest = Kanon {"Kanon-00_Title2","2838689"}

    GOLD = []Kanon {
        KossonGOLD001,KossonGOLD002,KossonGOLD003,KossonGOLD004,KossonGOLD005,KossonGOLD006,KossonGOLD007,KossonGOLD008,KossonGOLD009,KossonGOLD010,
        KossonGOLD011,KossonGOLD012,KossonGOLD013,KossonGOLD014,KossonGOLD015,KossonGOLD016,KossonGOLD017,KossonGOLD018,KossonGOLD019,KossonGOLD020,
        KossonGOLD021,KossonGOLD022,KossonGOLD023,KossonGOLD024,KossonGOLD025,KossonGOLD026,KossonGOLD027,KossonGOLD028,KossonGOLD029,KossonGOLD030,
        KossonGOLD031,KossonGOLD032,KossonGOLD033,KossonGOLD034,KossonGOLD035,KossonGOLD036,KossonGOLD037,KossonGOLD038,KossonGOLD039,KossonGOLD040,
        KossonGOLD041,KossonGOLD042,KossonGOLD043,KossonGOLD044,KossonGOLD045,KossonGOLD046,KossonGOLD047,KossonGOLD048,KossonGOLD049,KossonGOLD050,
        KossonGOLD051,KossonGOLD052,KossonGOLD053,KossonGOLD054,KossonGOLD055,KossonGOLD056,KossonGOLD057,KossonGOLD058,KossonGOLD059,KossonGOLD060,
        KossonGOLD061,KossonGOLD062,KossonGOLD063,KossonGOLD064,KossonGOLD065,KossonGOLD066,KossonGOLD067,KossonGOLD068,KossonGOLD069,KossonGOLD070,
        KossonGOLD071,KossonGOLD072,KossonGOLD073,KossonGOLD074,KossonGOLD075,KossonGOLD076,KossonGOLD077,KossonGOLD078,KossonGOLD079,KossonGOLD080,
        KossonGOLD081,KossonGOLD082,KossonGOLD083,KossonGOLD084,KossonGOLD085,KossonGOLD086,KossonGOLD087,KossonGOLD088,KossonGOLD089,KossonGOLD090,
        KossonGOLD091,KossonGOLD092,KossonGOLD093,KossonGOLD094,KossonGOLD095,KossonGOLD096,KossonGOLD097,KossonGOLD098,KossonGOLD099,KossonGOLD100,
    }

    SILVER = []Kanon {
        KossonSILVER0001,KossonSILVER0002,KossonSILVER0003,KossonSILVER0004,KossonSILVER0005,KossonSILVER0006,KossonSILVER0007,KossonSILVER0008,KossonSILVER0009,KossonSILVER0010,
        KossonSILVER0011,KossonSILVER0012,KossonSILVER0013,KossonSILVER0014,KossonSILVER0015,KossonSILVER0016,KossonSILVER0017,KossonSILVER0018,KossonSILVER0019,KossonSILVER0020,
        KossonSILVER0021,KossonSILVER0022,KossonSILVER0023,KossonSILVER0024,KossonSILVER0025,KossonSILVER0026,KossonSILVER0027,KossonSILVER0028,KossonSILVER0029,KossonSILVER0030,
        KossonSILVER0031,KossonSILVER0032,KossonSILVER0033,KossonSILVER0034,KossonSILVER0035,KossonSILVER0036,KossonSILVER0037,KossonSILVER0038,KossonSILVER0039,KossonSILVER0040,
        KossonSILVER0041,KossonSILVER0042,KossonSILVER0043,KossonSILVER0044,KossonSILVER0045,KossonSILVER0046,KossonSILVER0047,KossonSILVER0048,KossonSILVER0049,KossonSILVER0050,
        KossonSILVER0051,KossonSILVER0052,KossonSILVER0053,KossonSILVER0054,KossonSILVER0055,KossonSILVER0056,KossonSILVER0057,KossonSILVER0058,KossonSILVER0059,KossonSILVER0060,
        KossonSILVER0061,KossonSILVER0062,KossonSILVER0063,KossonSILVER0064,KossonSILVER0065,KossonSILVER0066,KossonSILVER0067,KossonSILVER0068,KossonSILVER0069,KossonSILVER0070,
        KossonSILVER0071,KossonSILVER0072,KossonSILVER0073,KossonSILVER0074,KossonSILVER0075,KossonSILVER0076,KossonSILVER0077,KossonSILVER0078,KossonSILVER0079,KossonSILVER0080,
        KossonSILVER0081,KossonSILVER0082,KossonSILVER0083,KossonSILVER0084,KossonSILVER0085,KossonSILVER0086,KossonSILVER0087,KossonSILVER0088,KossonSILVER0089,KossonSILVER0090,
        KossonSILVER0091,KossonSILVER0092,KossonSILVER0093,KossonSILVER0094,KossonSILVER0095,KossonSILVER0096,KossonSILVER0097,KossonSILVER0098,KossonSILVER0099,KossonSILVER0100,
        KossonSILVER0101,KossonSILVER0102,KossonSILVER0103,KossonSILVER0104,KossonSILVER0105,KossonSILVER0106,KossonSILVER0107,KossonSILVER0108,KossonSILVER0109,KossonSILVER0110,
        KossonSILVER0111,KossonSILVER0112,KossonSILVER0113,KossonSILVER0114,KossonSILVER0115,KossonSILVER0116,KossonSILVER0117,KossonSILVER0118,KossonSILVER0119,KossonSILVER0120,
        KossonSILVER0121,KossonSILVER0122,KossonSILVER0123,KossonSILVER0124,KossonSILVER0125,KossonSILVER0126,KossonSILVER0127,KossonSILVER0128,KossonSILVER0129,KossonSILVER0130,
        KossonSILVER0131,KossonSILVER0132,KossonSILVER0133,KossonSILVER0134,KossonSILVER0135,KossonSILVER0136,KossonSILVER0137,KossonSILVER0138,KossonSILVER0139,KossonSILVER0140,
        KossonSILVER0141,KossonSILVER0142,KossonSILVER0143,KossonSILVER0144,KossonSILVER0145,KossonSILVER0146,KossonSILVER0147,KossonSILVER0148,KossonSILVER0149,KossonSILVER0150,
        KossonSILVER0151,KossonSILVER0152,KossonSILVER0153,KossonSILVER0154,KossonSILVER0155,KossonSILVER0156,KossonSILVER0157,KossonSILVER0158,KossonSILVER0159,KossonSILVER0160,
        KossonSILVER0161,KossonSILVER0162,KossonSILVER0163,KossonSILVER0164,KossonSILVER0165,KossonSILVER0166,KossonSILVER0167,KossonSILVER0168,KossonSILVER0169,KossonSILVER0170,
        KossonSILVER0171,KossonSILVER0172,KossonSILVER0173,KossonSILVER0174,KossonSILVER0175,KossonSILVER0176,KossonSILVER0177,KossonSILVER0178,KossonSILVER0179,KossonSILVER0180,
        KossonSILVER0181,KossonSILVER0182,KossonSILVER0183,KossonSILVER0184,KossonSILVER0185,KossonSILVER0186,KossonSILVER0187,KossonSILVER0188,KossonSILVER0189,KossonSILVER0190,
        KossonSILVER0191,KossonSILVER0192,KossonSILVER0193,KossonSILVER0194,KossonSILVER0195,KossonSILVER0196,KossonSILVER0197,KossonSILVER0198,KossonSILVER0199,KossonSILVER0200,

        KossonSILVER0201,KossonSILVER0202,KossonSILVER0203,KossonSILVER0204,KossonSILVER0205,KossonSILVER0206,KossonSILVER0207,KossonSILVER0208,KossonSILVER0209,KossonSILVER0210,
        KossonSILVER0211,KossonSILVER0212,KossonSILVER0213,KossonSILVER0214,KossonSILVER0215,KossonSILVER0216,KossonSILVER0217,KossonSILVER0218,KossonSILVER0219,KossonSILVER0220,
        KossonSILVER0221,KossonSILVER0222,KossonSILVER0223,KossonSILVER0224,KossonSILVER0225,KossonSILVER0226,KossonSILVER0227,KossonSILVER0228,KossonSILVER0229,KossonSILVER0230,
        KossonSILVER0231,KossonSILVER0232,KossonSILVER0233,KossonSILVER0234,KossonSILVER0235,KossonSILVER0236,KossonSILVER0237,KossonSILVER0238,KossonSILVER0239,KossonSILVER0240,
        KossonSILVER0241,KossonSILVER0242,KossonSILVER0243,KossonSILVER0244,KossonSILVER0245,KossonSILVER0246,KossonSILVER0247,KossonSILVER0248,KossonSILVER0249,KossonSILVER0250,
        KossonSILVER0251,KossonSILVER0252,KossonSILVER0253,KossonSILVER0254,KossonSILVER0255,KossonSILVER0256,KossonSILVER0257,KossonSILVER0258,KossonSILVER0259,KossonSILVER0260,
        KossonSILVER0261,KossonSILVER0262,KossonSILVER0263,KossonSILVER0264,KossonSILVER0265,KossonSILVER0266,KossonSILVER0267,KossonSILVER0268,KossonSILVER0269,KossonSILVER0270,
        KossonSILVER0271,KossonSILVER0272,KossonSILVER0273,KossonSILVER0274,KossonSILVER0275,KossonSILVER0276,KossonSILVER0277,KossonSILVER0278,KossonSILVER0279,KossonSILVER0280,
        KossonSILVER0281,KossonSILVER0282,KossonSILVER0283,KossonSILVER0284,KossonSILVER0285,KossonSILVER0286,KossonSILVER0287,KossonSILVER0288,KossonSILVER0289,KossonSILVER0290,
        KossonSILVER0291,KossonSILVER0292,KossonSILVER0293,KossonSILVER0294,KossonSILVER0295,KossonSILVER0296,KossonSILVER0297,KossonSILVER0298,KossonSILVER0299,KossonSILVER0300,

        KossonSILVER0301,KossonSILVER0302,KossonSILVER0303,KossonSILVER0304,KossonSILVER0305,KossonSILVER0306,KossonSILVER0307,KossonSILVER0308,KossonSILVER0309,KossonSILVER0310,
        KossonSILVER0311,KossonSILVER0312,KossonSILVER0313,KossonSILVER0314,KossonSILVER0315,KossonSILVER0316,KossonSILVER0317,KossonSILVER0318,KossonSILVER0319,KossonSILVER0320,
        KossonSILVER0321,KossonSILVER0322,KossonSILVER0323,KossonSILVER0324,KossonSILVER0325,KossonSILVER0326,KossonSILVER0327,KossonSILVER0328,KossonSILVER0329,KossonSILVER0330,
        KossonSILVER0331,KossonSILVER0332,KossonSILVER0333,KossonSILVER0334,KossonSILVER0335,KossonSILVER0336,KossonSILVER0337,KossonSILVER0338,KossonSILVER0339,KossonSILVER0340,
        KossonSILVER0341,KossonSILVER0342,KossonSILVER0343,KossonSILVER0344,KossonSILVER0345,KossonSILVER0346,KossonSILVER0347,KossonSILVER0348,KossonSILVER0349,KossonSILVER0350,
        KossonSILVER0351,KossonSILVER0352,KossonSILVER0353,KossonSILVER0354,KossonSILVER0355,KossonSILVER0356,KossonSILVER0357,KossonSILVER0358,KossonSILVER0359,KossonSILVER0360,
        KossonSILVER0361,KossonSILVER0362,KossonSILVER0363,KossonSILVER0364,KossonSILVER0365,KossonSILVER0366,KossonSILVER0367,KossonSILVER0368,KossonSILVER0369,KossonSILVER0370,
        KossonSILVER0371,KossonSILVER0372,KossonSILVER0373,KossonSILVER0374,KossonSILVER0375,KossonSILVER0376,KossonSILVER0377,KossonSILVER0378,KossonSILVER0379,KossonSILVER0380,
        KossonSILVER0381,KossonSILVER0382,KossonSILVER0383,KossonSILVER0384,KossonSILVER0385,KossonSILVER0386,KossonSILVER0387,KossonSILVER0388,KossonSILVER0389,KossonSILVER0390,
        KossonSILVER0391,KossonSILVER0392,KossonSILVER0393,KossonSILVER0394,KossonSILVER0395,KossonSILVER0396,KossonSILVER0397,KossonSILVER0398,KossonSILVER0399,KossonSILVER0400,

        KossonSILVER0401,KossonSILVER0402,KossonSILVER0403,KossonSILVER0404,KossonSILVER0405,KossonSILVER0406,KossonSILVER0407,KossonSILVER0408,KossonSILVER0409,KossonSILVER0410,
        KossonSILVER0411,KossonSILVER0412,KossonSILVER0413,KossonSILVER0414,KossonSILVER0415,KossonSILVER0416,KossonSILVER0417,KossonSILVER0418,KossonSILVER0419,KossonSILVER0420,
        KossonSILVER0421,KossonSILVER0422,KossonSILVER0423,KossonSILVER0424,KossonSILVER0425,KossonSILVER0426,KossonSILVER0427,KossonSILVER0428,KossonSILVER0429,KossonSILVER0430,
        KossonSILVER0431,KossonSILVER0432,KossonSILVER0433,KossonSILVER0434,KossonSILVER0435,KossonSILVER0436,KossonSILVER0437,KossonSILVER0438,KossonSILVER0439,KossonSILVER0440,
        KossonSILVER0441,KossonSILVER0442,KossonSILVER0443,KossonSILVER0444,KossonSILVER0445,KossonSILVER0446,KossonSILVER0447,KossonSILVER0448,KossonSILVER0449,KossonSILVER0450,
        KossonSILVER0451,KossonSILVER0452,KossonSILVER0453,KossonSILVER0454,KossonSILVER0455,KossonSILVER0456,KossonSILVER0457,KossonSILVER0458,KossonSILVER0459,KossonSILVER0460,
        KossonSILVER0461,KossonSILVER0462,KossonSILVER0463,KossonSILVER0464,KossonSILVER0465,KossonSILVER0466,KossonSILVER0467,KossonSILVER0468,KossonSILVER0469,KossonSILVER0470,
        KossonSILVER0471,KossonSILVER0472,KossonSILVER0473,KossonSILVER0474,KossonSILVER0475,KossonSILVER0476,KossonSILVER0477,KossonSILVER0478,KossonSILVER0479,KossonSILVER0480,
        KossonSILVER0481,KossonSILVER0482,KossonSILVER0483,KossonSILVER0484,KossonSILVER0485,KossonSILVER0486,KossonSILVER0487,KossonSILVER0488,KossonSILVER0489,KossonSILVER0490,
        KossonSILVER0491,KossonSILVER0492,KossonSILVER0493,KossonSILVER0494,KossonSILVER0495,KossonSILVER0496,KossonSILVER0497,KossonSILVER0498,KossonSILVER0499,KossonSILVER0500,

        KossonSILVER0501,KossonSILVER0502,KossonSILVER0503,KossonSILVER0504,KossonSILVER0505,KossonSILVER0506,KossonSILVER0507,KossonSILVER0508,KossonSILVER0509,KossonSILVER0510,
        KossonSILVER0511,KossonSILVER0512,KossonSILVER0513,KossonSILVER0514,KossonSILVER0515,KossonSILVER0516,KossonSILVER0517,KossonSILVER0518,KossonSILVER0519,KossonSILVER0520,
        KossonSILVER0521,KossonSILVER0522,KossonSILVER0523,KossonSILVER0524,KossonSILVER0525,KossonSILVER0526,KossonSILVER0527,KossonSILVER0528,KossonSILVER0529,KossonSILVER0530,
        KossonSILVER0531,KossonSILVER0532,KossonSILVER0533,KossonSILVER0534,KossonSILVER0535,KossonSILVER0536,KossonSILVER0537,KossonSILVER0538,KossonSILVER0539,KossonSILVER0540,
        KossonSILVER0541,KossonSILVER0542,KossonSILVER0543,KossonSILVER0544,KossonSILVER0545,KossonSILVER0546,KossonSILVER0547,KossonSILVER0548,KossonSILVER0549,KossonSILVER0550,
        KossonSILVER0551,KossonSILVER0552,KossonSILVER0553,KossonSILVER0554,KossonSILVER0555,KossonSILVER0556,KossonSILVER0557,KossonSILVER0558,KossonSILVER0559,KossonSILVER0560,
        KossonSILVER0561,KossonSILVER0562,KossonSILVER0563,KossonSILVER0564,KossonSILVER0565,KossonSILVER0566,KossonSILVER0567,KossonSILVER0568,KossonSILVER0569,KossonSILVER0570,
        KossonSILVER0571,KossonSILVER0572,KossonSILVER0573,KossonSILVER0574,KossonSILVER0575,KossonSILVER0576,KossonSILVER0577,KossonSILVER0578,KossonSILVER0579,KossonSILVER0580,
        KossonSILVER0581,KossonSILVER0582,KossonSILVER0583,KossonSILVER0584,KossonSILVER0585,KossonSILVER0586,KossonSILVER0587,KossonSILVER0588,KossonSILVER0589,KossonSILVER0590,
        KossonSILVER0591,KossonSILVER0592,KossonSILVER0593,KossonSILVER0594,KossonSILVER0595,KossonSILVER0596,KossonSILVER0597,KossonSILVER0598,KossonSILVER0599,KossonSILVER0600,

        KossonSILVER0601,KossonSILVER0602,KossonSILVER0603,KossonSILVER0604,KossonSILVER0605,KossonSILVER0606,KossonSILVER0607,KossonSILVER0608,KossonSILVER0609,KossonSILVER0610,
        KossonSILVER0611,KossonSILVER0612,KossonSILVER0613,KossonSILVER0614,KossonSILVER0615,KossonSILVER0616,KossonSILVER0617,KossonSILVER0618,KossonSILVER0619,KossonSILVER0620,
        KossonSILVER0621,KossonSILVER0622,KossonSILVER0623,KossonSILVER0624,KossonSILVER0625,KossonSILVER0626,KossonSILVER0627,KossonSILVER0628,KossonSILVER0629,KossonSILVER0630,
        KossonSILVER0631,KossonSILVER0632,KossonSILVER0633,KossonSILVER0634,KossonSILVER0635,KossonSILVER0636,KossonSILVER0637,KossonSILVER0638,KossonSILVER0639,KossonSILVER0640,
        KossonSILVER0641,KossonSILVER0642,KossonSILVER0643,KossonSILVER0644,KossonSILVER0645,KossonSILVER0646,KossonSILVER0647,KossonSILVER0648,KossonSILVER0649,KossonSILVER0650,
        KossonSILVER0651,KossonSILVER0652,KossonSILVER0653,KossonSILVER0654,KossonSILVER0655,KossonSILVER0656,KossonSILVER0657,KossonSILVER0658,KossonSILVER0659,KossonSILVER0660,
        KossonSILVER0661,KossonSILVER0662,KossonSILVER0663,KossonSILVER0664,KossonSILVER0665,KossonSILVER0666,KossonSILVER0667,KossonSILVER0668,KossonSILVER0669,KossonSILVER0670,
        KossonSILVER0671,KossonSILVER0672,KossonSILVER0673,KossonSILVER0674,KossonSILVER0675,KossonSILVER0676,KossonSILVER0677,KossonSILVER0678,KossonSILVER0679,KossonSILVER0680,
        KossonSILVER0681,KossonSILVER0682,KossonSILVER0683,KossonSILVER0684,KossonSILVER0685,KossonSILVER0686,KossonSILVER0687,KossonSILVER0688,KossonSILVER0689,KossonSILVER0690,
        KossonSILVER0691,KossonSILVER0692,KossonSILVER0693,KossonSILVER0694,KossonSILVER0695,KossonSILVER0696,KossonSILVER0697,KossonSILVER0698,KossonSILVER0699,KossonSILVER0700,

        KossonSILVER0701,KossonSILVER0702,KossonSILVER0703,KossonSILVER0704,KossonSILVER0705,KossonSILVER0706,KossonSILVER0707,KossonSILVER0708,KossonSILVER0709,KossonSILVER0710,
        KossonSILVER0711,KossonSILVER0712,KossonSILVER0713,KossonSILVER0714,KossonSILVER0715,KossonSILVER0716,KossonSILVER0717,KossonSILVER0718,KossonSILVER0719,KossonSILVER0720,
        KossonSILVER0721,KossonSILVER0722,KossonSILVER0723,KossonSILVER0724,KossonSILVER0725,KossonSILVER0726,KossonSILVER0727,KossonSILVER0728,KossonSILVER0729,KossonSILVER0730,
        KossonSILVER0731,KossonSILVER0732,KossonSILVER0733,KossonSILVER0734,KossonSILVER0735,KossonSILVER0736,KossonSILVER0737,KossonSILVER0738,KossonSILVER0739,KossonSILVER0740,
        KossonSILVER0741,KossonSILVER0742,KossonSILVER0743,KossonSILVER0744,KossonSILVER0745,KossonSILVER0746,KossonSILVER0747,KossonSILVER0748,KossonSILVER0749,KossonSILVER0750,
        KossonSILVER0751,KossonSILVER0752,KossonSILVER0753,KossonSILVER0754,KossonSILVER0755,KossonSILVER0756,KossonSILVER0757,KossonSILVER0758,KossonSILVER0759,KossonSILVER0760,
        KossonSILVER0761,KossonSILVER0762,KossonSILVER0763,KossonSILVER0764,KossonSILVER0765,KossonSILVER0766,KossonSILVER0767,KossonSILVER0768,KossonSILVER0769,KossonSILVER0770,
        KossonSILVER0771,KossonSILVER0772,KossonSILVER0773,KossonSILVER0774,KossonSILVER0775,KossonSILVER0776,KossonSILVER0777,KossonSILVER0778,KossonSILVER0779,KossonSILVER0780,
        KossonSILVER0781,KossonSILVER0782,KossonSILVER0783,KossonSILVER0784,KossonSILVER0785,KossonSILVER0786,KossonSILVER0787,KossonSILVER0788,KossonSILVER0789,KossonSILVER0790,
        KossonSILVER0791,KossonSILVER0792,KossonSILVER0793,KossonSILVER0794,KossonSILVER0795,KossonSILVER0796,KossonSILVER0797,KossonSILVER0798,KossonSILVER0799,KossonSILVER0800,

        KossonSILVER0801,KossonSILVER0802,KossonSILVER0803,KossonSILVER0804,KossonSILVER0805,KossonSILVER0806,KossonSILVER0807,KossonSILVER0808,KossonSILVER0809,KossonSILVER0810,
        KossonSILVER0811,KossonSILVER0812,KossonSILVER0813,KossonSILVER0814,KossonSILVER0815,KossonSILVER0816,KossonSILVER0817,KossonSILVER0818,KossonSILVER0819,KossonSILVER0820,
        KossonSILVER0821,KossonSILVER0822,KossonSILVER0823,KossonSILVER0824,KossonSILVER0825,KossonSILVER0826,KossonSILVER0827,KossonSILVER0828,KossonSILVER0829,KossonSILVER0830,
        KossonSILVER0831,KossonSILVER0832,KossonSILVER0833,KossonSILVER0834,KossonSILVER0835,KossonSILVER0836,KossonSILVER0837,KossonSILVER0838,KossonSILVER0839,KossonSILVER0840,
        KossonSILVER0841,KossonSILVER0842,KossonSILVER0843,KossonSILVER0844,KossonSILVER0845,KossonSILVER0846,KossonSILVER0847,KossonSILVER0848,KossonSILVER0849,KossonSILVER0850,
        KossonSILVER0851,KossonSILVER0852,KossonSILVER0853,KossonSILVER0854,KossonSILVER0855,KossonSILVER0856,KossonSILVER0857,KossonSILVER0858,KossonSILVER0859,KossonSILVER0860,
        KossonSILVER0861,KossonSILVER0862,KossonSILVER0863,KossonSILVER0864,KossonSILVER0865,KossonSILVER0866,KossonSILVER0867,KossonSILVER0868,KossonSILVER0869,KossonSILVER0870,
        KossonSILVER0871,KossonSILVER0872,KossonSILVER0873,KossonSILVER0874,KossonSILVER0875,KossonSILVER0876,KossonSILVER0877,KossonSILVER0878,KossonSILVER0879,KossonSILVER0880,
        KossonSILVER0881,KossonSILVER0882,KossonSILVER0883,KossonSILVER0884,KossonSILVER0885,KossonSILVER0886,KossonSILVER0887,KossonSILVER0888,KossonSILVER0889,KossonSILVER0890,
        KossonSILVER0891,KossonSILVER0892,KossonSILVER0893,KossonSILVER0894,KossonSILVER0895,KossonSILVER0896,KossonSILVER0897,KossonSILVER0898,KossonSILVER0899,KossonSILVER0900,

        KossonSILVER0901,KossonSILVER0902,KossonSILVER0903,KossonSILVER0904,KossonSILVER0905,KossonSILVER0906,KossonSILVER0907,KossonSILVER0908,KossonSILVER0909,KossonSILVER0910,
        KossonSILVER0911,KossonSILVER0912,KossonSILVER0913,KossonSILVER0914,KossonSILVER0915,KossonSILVER0916,KossonSILVER0917,KossonSILVER0918,KossonSILVER0919,KossonSILVER0920,
        KossonSILVER0921,KossonSILVER0922,KossonSILVER0923,KossonSILVER0924,KossonSILVER0925,KossonSILVER0926,KossonSILVER0927,KossonSILVER0928,KossonSILVER0929,KossonSILVER0930,
        KossonSILVER0931,KossonSILVER0932,KossonSILVER0933,KossonSILVER0934,KossonSILVER0935,KossonSILVER0936,KossonSILVER0937,KossonSILVER0938,KossonSILVER0939,KossonSILVER0940,
        KossonSILVER0941,KossonSILVER0942,KossonSILVER0943,KossonSILVER0944,KossonSILVER0945,KossonSILVER0946,KossonSILVER0947,KossonSILVER0948,KossonSILVER0949,KossonSILVER0950,
        KossonSILVER0951,KossonSILVER0952,KossonSILVER0953,KossonSILVER0954,KossonSILVER0955,KossonSILVER0956,KossonSILVER0957,KossonSILVER0958,KossonSILVER0959,KossonSILVER0960,
        KossonSILVER0961,KossonSILVER0962,KossonSILVER0963,KossonSILVER0964,KossonSILVER0965,KossonSILVER0966,KossonSILVER0967,KossonSILVER0968,KossonSILVER0969,KossonSILVER0970,
        KossonSILVER0971,KossonSILVER0972,KossonSILVER0973,KossonSILVER0974,KossonSILVER0975,KossonSILVER0976,KossonSILVER0977,KossonSILVER0978,KossonSILVER0979,KossonSILVER0980,
        KossonSILVER0981,KossonSILVER0982,KossonSILVER0983,KossonSILVER0984,KossonSILVER0985,KossonSILVER0986,KossonSILVER0987,KossonSILVER0988,KossonSILVER0989,KossonSILVER0990,
        KossonSILVER0991,KossonSILVER0992,KossonSILVER0993,KossonSILVER0994,KossonSILVER0995,KossonSILVER0996,KossonSILVER0997,KossonSILVER0998,KossonSILVER0999,KossonSILVER1000,
    }



    KossonGOLD001 = Kanon {"navel of life","the navel of life OM"}

    //Kanon 01
    KossonGOLD002 = Kanon {"Kanon  1.   1", "Im Namen der Schöpfung, die da ist Liebe, Wissen, Weisheit, Wahrheit und (relative) Vollkommenheit."}
    KossonGOLD003 = Kanon {"Kanon  1.   2", "Preis (Ehre, Dank) sei ihr, der Schöpfung aller Kreationen."}
    KossonGOLD004 = Kanon {"Kanon  1.   3", "Ihr allein lebet der Mensch, und ihr allein leben wir stetig."}
    KossonGOLD005 = Kanon {"Kanon  1.   4", "Ihre Gesetze und Gebote allein befolgen wir."}
    KossonGOLD006 = Kanon {"Kanon  1.   5", "Die Schöpfung selbst führet uns durch ihre Gesetze und Gebote zum Ziel der Evolution."}
    KossonGOLD007 = Kanon {"Kanon  1.   6", "Sie führet uns auf dem Wege derer, die den Pfad der Wahrheit und Liebe beschreiten und die den Weg der Wahrheit schon vor uns beschritten und kein Missfallen erreget haben noch irregegangen sind."}
    KossonGOLD008 = Kanon {"Kanon  1.   7", "Im Namen der Schöpfung, das ist unser Wille und unser Ziel der Evolution."}

    //Kanon 02
    KossonGOLD009 = Kanon {"Kanon  2.   1", "Im Namen der Schöpfung, der Liebevollen, der Weisen."}
    KossonGOLD010 = Kanon {"Kanon  2.   2", "So sprichet der JHWH und der Prophet, die da sind Weisheitkönig über die irdischen Menschengeschlechter:"}
    KossonGOLD011 = Kanon {"Kanon  2.   3", "Der JHWH und der Prophet offenbaren die Gesetze und Gebote der Schöpfung, die da ist die einzige Schöpfung ihres Himmels (Universums)."}
    KossonGOLD012 = Kanon {"Kanon  2.   4", "Preis (Ehre) sei ihr, die da ist das Leben und das SEIN (Schöpfungsleben)."}
    KossonGOLD013 = Kanon {"Kanon  2.   5", "Ihre siebenheitlichen Gesetzmässigkeiten offenkunden der JHWH und der Prophet, die da sind Mensch und kundige Offenbarung alles Wahrheit, aller Liebe, alle Weisheit und des Wissens und der Wege der Schöpfung."}
    KossonGOLD014 = Kanon {"Kanon  2.   6", "Sie jedoch sind nicht eine Vertretung der Schöpfung, sondern Erkenntnis, Offenbarung und Pflichterfüllung der Wahrheit und Liebe in freier Form, zur Befolgung der Gesetze und Gebote der Schöpfung."}
    KossonGOLD015 = Kanon {"Kanon  2.   7", "Und so sprichet der JHWH und der Prophet in Weisheit und freier Pflichterkennung und Pflichterfüllung für die Erfüllung der Gesetze und Gebote der Schöpfung:"}
    KossonGOLD016 = Kanon {"Kanon  2.   8", "Nicht ist verehrungswürdig (und anbetungswürdig) der JHWH, denn dies gebühret allein der Schöpfung."}
    KossonGOLD017 = Kanon {"Kanon  2.   9", "Und nicht ist verehrungswürdig (vergötterungswürdig) und anbetungswürdig (ergebungswürdig) der Prophet, denn dies gebühret allein der Schöpfung."}
    KossonGOLD018 = Kanon {"Kanon  2.  10", "Und nicht ist verehrungswürdig (vergötterungswürdig) und anbetungswürdig (ergebungswürdig) die Geistführer des Propheten, denn dies gebühret allein der Schöpfung."}
    KossonGOLD019 = Kanon {"Kanon  2.  11", "Niemals verrichte ein Mensch sein Gebet (Anbettelung,Anflehung) and den JHWH oder an den Propheten oder an des Propheten Geistführer, denn Gebetsverrichtung (Insichgehen in Gedanken) gebühret allein an den Geist (Bewusstsein)."}
    KossonGOLD020 = Kanon {"Kanon  2.  12", "Niemals verehre ein Mensch den JHWH oder den Propheten oder des Propheten Geistführer, den Verehrung (Ehrwürdigung) gebühret allein der Schöpfung."}
    KossonGOLD021 = Kanon {"Kanon  2.  13", "Verehrung und Anbetung gebühret allein der Schöpfung, die da ist das Höchste und die (relativ) vollkommenste Vollkommenheit in ihrem Himmel(Universum)."}
    KossonGOLD022 = Kanon {"Kanon  2.  14", "Achtet, übet und bringet jedoch Respekt und Ehrfurcht (Ehrung, Ehrwürdigung, Ehrerbietung, Ehrwürdigkeit) entgegen allem Leben, so auch dem JHWH der irdischen Menschengeschlechter und auch dem Propheten und des Propheten Geistführer."}
    KossonGOLD023 = Kanon {"Kanon  2.  15", "Achtet, übet und bringet jedoch Respekt und Ehrfurcht (Ehrung, Ehrerbietung, Ehrwürdigung, Ehrwürdigkeit) entgegen allem Leben."}
    KossonGOLD024 = Kanon {"Kanon  2.  16", "Achtet, übet und bringet jedoch Respekt und Ehrfurcht (Ehrung, Ehrerbietung, Ehrwürdigung, Ehrwürdigkeit) entgegen dem JHWH, der da ist der Weisheitskönig der Menschengeschlechter dreier Welten."}
    KossonGOLD025 = Kanon {"Kanon  2.  17", "Achtet, übet und bringet jedoch Respekt und Ehrfurcht (Ehrung, Ehrerbietung, Ehrwürdigung, Ehrwürdigkeit) entgegen dem Propheten, der da ist wahrlicher Prophet für die irsichen Menschengeschlechter in Wahrheit, Liebe, Weisheit und Wissen."}
    KossonGOLD026 = Kanon {"Kanon  2.  18", "Achtet, übet und bringet jedoch Respekt und Ehrfurcht (Ehrung, Ehrerbietung, Ehrwürdigung, Ehrwürdigkeit) entgegen dem Propheten, der da ist Mittler, Künder, Wissender und Liebevoller und Weiser zwischen den zu belehrenden irdischen Menschengeschlechter und der Lehre der Wahrheit."}
    KossonGOLD027 = Kanon {"Kanon  2.  19", "Achtet, übet und bringet jedoch Respekt und Ehrfurcht (Ehrung, Ehrerbietung, Ehrwürdigung, Ehrwürdigkeit) entgegen den Geistführern des Propheten, die da sind die Wissenden und Belehrer der unwissenden irdischen Menschengeschlechter im Auftrage des Propheten."}
    KossonGOLD028 = Kanon {"Kanon  2.  20", "Befolget die Lehre der Gesetze und Gebote der Schöpfung, denn allein ihre Befolgung gewähret das wahrliche Leben und die Evolution zum SEIN (Schöpfungsleben)."}
    KossonGOLD029 = Kanon {"Kanon  2.  21", "In Namen der Schöpfung, der Liebevollen, der Weisen."}

    //Kanon 03
    KossonGOLD030 = Kanon {"Kanon  3.   1", "Im Namen der Schöpfung, der Wissenden, der Kreierenden."}
    KossonGOLD031 = Kanon {"Kanon  3.   2", "So sprichet der JHWH, der da ist der JHWH der irdischen Menschengeschlechter:"}
    KossonGOLD032 = Kanon {"Kanon  3.   3", "Das ist das Buch der Gesetze und Gebote der Schöpfung, ausgelegt und erkläret mit verständlichem Wort."}
    KossonGOLD033 = Kanon {"Kanon  3.   4", "Das ist das Buch der Kenntnis der schöpferischen Gesetze und Gebote, wie es im Wort und Sinn ist gegeben durch den JHWH und den Propheten der irdischen Menschengeschlechter."}
    KossonGOLD034 = Kanon {"Kanon  3.   5", "Nicht ist der JHWH die Shöpfung selbst, sondern er ist Mensch und Weisheitskönig über die Menschengeschlechter der Erde in nichthierarchischer Form."}
    KossonGOLD035 = Kanon {"Kanon  3.   6", "Der JHWH ist der alleinige JHWH für die irdischen Menschengeschlechter und für die Weithergereisten aus den Tiefen des Himmels (Universums), die da sind seine Helfer und Wächter; doch ausser der JHWH ist noch der Prophet für die irdischen Menschengeschlechter, der da ist eine Kraft gleichsam dem JHWH."}
    KossonGOLD036 = Kanon {"Kanon  3.   7", "Dem JHWH und dem Propheten der irdischen Menschengeschlechter gebühret der Respekt un die Ehrfurcht (Ehrung, Ehrerbietung, Ehrwürdigung) der Menschen der Erde."}
    KossonGOLD037 = Kanon {"Kanon  3.   8", "Der Schöpfung aber gebühret die Ehre, der Respekt und die Ehrfurcht (Ehrung, Ehrerbietung, Ehrwürdigung) des JHWH, des Propheten und der Menschen."}
    KossonGOLD038 = Kanon {"Kanon  3.   9", "Die Schöpfung allein ist die Kraft aller Kreation, des Lebens und des SEINs (Schöpfungslebens)."}
    KossonGOLD039 = Kanon {"Kanon  3.  10", "Und nichts ist verehrungswürdig und anbentungswürdig ausser der Schöpfung."}
    KossonGOLD040 = Kanon {"Kanon  3.  11", "Der JHWH ist weise, doch er ist Mensch, und er ist König des Wissens, der Liebe, der Wahrheit, des Könnens und der Weisheit, so er also ist Weisheitskönig."}
    KossonGOLD041 = Kanon {"Kanon  3.  12", "Ausser dem JHWH ist nichts in gleicher Form des Menschen."}
    KossonGOLD042 = Kanon {"Kanon  3.  13", "Doch der JHWH ist Mensch und bleibet Mensch; bis er zu Geistform wandelt."}
    KossonGOLD043 = Kanon {"Kanon  3.  14", "Der JHWH ist Mensch, und über ihm thronet und schwebet unmessbar hoch die Schöpfung in Allmacht aller Schöpfung, und ihr ist in Kraft, Wissen, Liebe und Weisheit nichts Vergleichbares im Himmel(Universum)."}
    KossonGOLD044 = Kanon {"Kanon  3.  15", "Die Schöpfung ist die Schöpfung, und sie kreierete alle Schöpfung, die da ist existent in ihrem Himmel(Universum)."}
    KossonGOLD045 = Kanon {"Kanon  3.  16", "Die Schöpfung ist selbst Schöpfung, und sie kreierete alle Schöpfung, so das Leben, das SEIN (Schöpfungsleben), die Himmel (Universum, Ebenen und Dimensionen) und die Erden, die Galaxien und die Festen, und also jegliche Kreatur und alle Heere des Lebens."}
    KossonGOLD046 = Kanon {"Kanon  3.  17", "Die Schöpfung ist Schöpfung, und ausser der Schöpfung ist keine Schöpfung in ihrem Himmen(Universum)."}
    KossonGOLD047 = Kanon {"Kanon  3.  18", "Und die Schöpfung ist om."}
    KossonGOLD048 = Kanon {"Kanon  3.  19", "Der JHWH und der Prophet sind Mensch, und gewandt in Gesetz und Gebot der Schöpfung, und sie bringen den zu belehrenden Menschengeschlechter dar das Buch der Gesetze und Gebote der Schöpfung, denn sie sind Mensch und JHWH und Prophet und als solsche dafür in Pflicht."}
    KossonGOLD049 = Kanon {"Kanon  3.  20", "Über dem JHWH aber und über den Propheten herrschet und schwebet in unmessbarer Höhe die Schöpfung als endlose und allgrosszeitliche Kraft und als wahrliches SEIN (Schöpfungsleben)."}
    KossonGOLD050 = Kanon {"Kanon  3.  21", "Die Schöpfung ist die Schöpfung, und ausser ihr ist keine Schöpfung in ihrem Himmel(Universum)."}
    //Kanon 04
    KossonGOLD051 = Kanon {"Kanon  4.   1", "In Namen der Schöpfung, der Gnadvollen, der Alliebenden, der Gerechten, der Weisen, der Wissenden, der Allbarmherzigen."}
    KossonGOLD052 = Kanon {"Kanon  4.   2", "Preis (Ehre) sei der Schöpfung, die da ist die wahrliche Schöpfung und die einzige Schöpfung in ihrem Himmel(Universum)."}
    KossonGOLD053 = Kanon {"Kanon  4.   3", "Preis (Ehre) sei der Schöpfung, die da ist  die Schöpfung aller Kreationen und die da sind der Himmel(Universum), die Galaxien, die Gestirne und die Erden, die Himmel (Ebenen und Dimensionen), das Licht und die Dunkelheit, die Zeit und der Raum und alle Heere der lebendigen Formen des Lebens, jegliches nach seiner Art."}
    KossonGOLD054 = Kanon {"Kanon  4.   4", "Preis (Ehre) sei der Schöpfung, die da ist die einzige Schöpfung in ihrem Himmel(Universum) und die da ist die Schöpfung aller Kreationen in Raum und Zeit."}
    KossonGOLD055 = Kanon {"Kanon  4.   5", "Preis (Ehre) sei der Schöpfung, die da ist Gerechtigkeit."}
    KossonGOLD056 = Kanon {"Kanon  4.   6", "Preis (Ehre) sei der Schöpfung, die da ist Liebe."}
    KossonGOLD057 = Kanon {"Kanon  4.   7", "Preis (Ehre) sei der Schöpfung, die da ist Kraft."}
    KossonGOLD058 = Kanon {"Kanon  4.   8", "Preis (Ehre) sei der Schöpfung, die da ist Weisheit."}
    KossonGOLD059 = Kanon {"Kanon  4.   9", "Preis (Ehre) sei der Schöpfung, die da ist Wissen."}
    KossonGOLD060 = Kanon {"Kanon  4.  10", "Preis (Ehre) sei der Schöpfung, die da ist Barmherzigkeit."}
    KossonGOLD061 = Kanon {"Kanon  4.  11", "Preis (Ehre) sei der Schöpfung, die da ist Freiheit."}
    KossonGOLD062 = Kanon {"Kanon  4.  12", "Preis (Ehre) sei der Schöpfung, die da ist Gnade."}
    KossonGOLD063 = Kanon {"Kanon  4.  13", "Preis (Ehre) sei der Schöpfung, die da ist Gesetz."}
    KossonGOLD064 = Kanon {"Kanon  4.  14", "Preis (Ehre) sei der Schöpfung, die da ist Gebot."}
    KossonGOLD065 = Kanon {"Kanon  4.  15", "Preis (Ehre) sei der Schöpfung, die da ist Bündnis."}
    KossonGOLD066 = Kanon {"Kanon  4.  16", "Preis (Ehre) sei der Schöpfung, die da ist Erfüllung."}
    KossonGOLD067 = Kanon {"Kanon  4.  17", "Preis (Ehre) sei der Schöpfung, die da ist Evolution."}
    KossonGOLD068 = Kanon {"Kanon  4.  18", "Preis (Ehre) sei der Schöpfung, die da ist Leben."}
    KossonGOLD069 = Kanon {"Kanon  4.  19", "Preis (Ehre) sei der Schöpfung, die da ist Beistand."}
    KossonGOLD070 = Kanon {"Kanon  4.  20", "Preis (Ehre) sei der Schöpfung, die da ist Freude."}
    KossonGOLD071 = Kanon {"Kanon  4.  21", "Preis (Ehre) sei der Schöpfung, die da ist Schönheit."}
    KossonGOLD072 = Kanon {"Kanon  4.  22", "Preis (Ehre) sei der Schöpfung, die da ist Frieden."}
    KossonGOLD073 = Kanon {"Kanon  4.  23", "Preis (Ehre) sei der Schöpfung, die da ist Unfehlbarkeit."}
    KossonGOLD074 = Kanon {"Kanon  4.  24", "Preis (Ehre) sei der Schöpfung, die da ist Ausgeglichenheit."}
    KossonGOLD075 = Kanon {"Kanon  4.  25", "Preis (Ehre) sei der Schöpfung, die da ist Geist."}
    KossonGOLD076 = Kanon {"Kanon  4.  26", "Preis (Ehre) sei der Schöpfung, die da ist Allzeitlichkeit."}
    KossonGOLD077 = Kanon {"Kanon  4.  27", "Preis (Ehre) sei der Schöpfung, die da ist Folgerichtigkeit."}
    KossonGOLD078 = Kanon {"Kanon  4.  28", "Preis (Ehre) sei der Schöpfung, die da ist Werden."}
    KossonGOLD079 = Kanon {"Kanon  4.  29", "Preis (Ehre) sei der Schöpfung, die da ist Vollkommenheit."}
    KossonGOLD080 = Kanon {"Kanon  4.  30", "Preis (Ehre) sei der Schöpfung, die da ist Zufriedenheit."}
    KossonGOLD081 = Kanon {"Kanon  4.  31", "Preis (Ehre) sei der Schöpfung, die da ist Unerschöpflichkeit."}
    KossonGOLD082 = Kanon {"Kanon  4.  32", "Preis (Ehre) sei der Schöpfung, die da ist Allmacht."}
    KossonGOLD083 = Kanon {"Kanon  4.  33", "Preis (Ehre) sei der Schöpfung, die da ist Lieblichkeit."}
    KossonGOLD084 = Kanon {"Kanon  4.  34", "Preis (Ehre) sei der Schöpfung, die da ist Unendlichkeit."}
    KossonGOLD085 = Kanon {"Kanon  4.  35", "Preis (Ehre) sei der Schöpfung, die da ist Verbundenheit."}
    KossonGOLD086 = Kanon {"Kanon  4.  36", "Preis (Ehre) sei der Schöpfung, die da ist Empfindung."}
    KossonGOLD087 = Kanon {"Kanon  4.  37", "Preis (Ehre) sei der Schöpfung, die da ist Erhörung."}
    KossonGOLD088 = Kanon {"Kanon  4.  38", "Preis (Ehre) sei der Schöpfung, die da ist Erhöhung."}
    KossonGOLD089 = Kanon {"Kanon  4.  39", "Preis (Ehre) sei der Schöpfung, die da ist SOHAR"}
    KossonGOLD090 = Kanon {"Kanon  4.  40", "Preis (Ehre) sei der Schöpfung, die da ist Sanftheit."}
    KossonGOLD091 = Kanon {"Kanon  4.  41", "Preis (Ehre) sei der Schöpfung, die da ist Klarheit."}
    KossonGOLD092 = Kanon {"Kanon  4.  42", "Preis (Ehre) sei der Schöpfung, die da ist Reinheit."}
    KossonGOLD093 = Kanon {"Kanon  4.  43", "Preis (Ehre) sei der Schöpfung, die da ist Wandlung."}
    KossonGOLD094 = Kanon {"Kanon  4.  44", "Preis (Ehre) sei der Schöpfung, die da ist Entstehung."}
    KossonGOLD095 = Kanon {"Kanon  4.  45", "Preis (Ehre) sei der Schöpfung, die da ist Zukunft."}
    KossonGOLD096 = Kanon {"Kanon  4.  46", "Preis (Ehre) sei der Schöpfung, die da ist Macht."}
    KossonGOLD097 = Kanon {"Kanon  4.  47", "Preis (Ehre) sei der Schöpfung, die da ist Ehrfurcht (Ehrung, Ehrerbietung, Erwürdigung, Ehrwürdigkeit)."}
    KossonGOLD098 = Kanon {"Kanon  4.  48", "Preis (Ehre) sei der Schöpfung, die da ist Allheit."}
    KossonGOLD099 = Kanon {"Kanon  4.  49", "Preis (Ehre) sei der Schöpfung, die da ist SEIN (Schöpfungsleben)."}

    KossonGOLD100 = Kanon {"navel of life","omfalon murado OM"}
)

var (
    //Kanon 5
    KossonSILVER0001 = Kanon {"Kanon  5.   1", "Im Namen der Schöpfung, die da ist die Schöpfung und wahrlich die Alleinige."}
    KossonSILVER0002 = Kanon {"Kanon  5.   2", "Der JHWH wachet als wahrlicher Weisheitskonig uber die Menschengeschlechter der Erde, und über ihm herrschet und schwebet die Schöpfung, deren Gesetze und Gebote er befolget in strengster Pflicht."}
    KossonSILVER0003 = Kanon {"Kanon  5.   3", "Der JHWH wachet in Belehrung, niemals jedoch als Herr, so er der Schöpfung gleichhandelt."}
    KossonSILVER0004 = Kanon {"Kanon  5.   4", "Der JHWH ist Weisheitskönig und wachend als solcher, nicht jedoch herrschend als Herr."}
    KossonSILVER0005 = Kanon {"Kanon  5.   5", "Die Schöpfung ist die Schöpfung, und keine Schöpfung ist in ihrem Himmel (Universum) neben ihr."}
    KossonSILVER0006 = Kanon {"Kanon  5.   6", "Der Mensch, der JHWH ist, ist Mensch, und als solcher ist er pflichtig für die Menschengeschlechter dreier Erden, denen er als Weisheitskönig obstehet."}
    KossonSILVER0007 = Kanon {"Kanon  5.   7", "Ausser dem JHWH der Erde, über dem die Schöpfung in unmessbarer Grösse schwebet, soll der Mensch keinen andern JHWH oder gar Götzen und Götter neben ihm haben."}
    KossonSILVER0008 = Kanon {"Kanon  5.   8", "Siehe also: Über die irdischen Menschengeschlechter wachet der Weisheitskönig, und in Belehrung der Prophet also."}
    KossonSILVER0009 = Kanon {"Kanon  5.   9", "Über dem Weisheitskönig und dem Propheten, die Mensch sind, schwebet aber in unermesslicher Höhe die Schöpfung, die als höchste und kreierende Kraft das SEIN (Schöpfungsleben) ist."}
    KossonSILVER0010 = Kanon {"Kanon  5.  10", "Der JHWH und der Prophet sind die Gesetzbringer und Gebotbringer für die Menschengeschlechter, und ihre Wünsche und Ratgebungen sollen befolget sein von Mann und Weib also, denn sie entstammen den Gesetzen und Geboten der Schöpfung, die stetige Befolgung finden müssen."}
    KossonSILVER0011 = Kanon {"Kanon  5.  11", "Der JHWH und der Prophet sind grossmütig in ihrer Liebe und in allem ihrem Tun und in ihrer Geduld also."}
    KossonSILVER0012 = Kanon {"Kanon  5.  12", "Der Mensch missachte nicht die Lehre des JHWH und des Propheten, denn sie beinhaltet die wahrlichen Erkenntnisse, Auslegungen und Ratgebungen der siebenheitlichen Ordnung, die da entspringen aus den Gesetzen und Geboten der Schöpfung."}
    KossonSILVER0013 = Kanon {"Kanon  5.  13", "Der JHWH und der Prophet erfassen und erkennen im gesamten die Gesetze und Gebote der Schöpfung in eigener Pflicht, und sie legen sie erklärend und ratgebend aus, denn sie sind Weisheitskönig und wissend und weise in allen Dingen."}
    KossonSILVER0014 = Kanon {"Kanon  5.  14", "Die Schöpfung stehet aber unmessbar viel höher als der JHWH und als der Prophet, die da Mensch sind, denn allein die Schöpfung ist das unmessbare Geheimnis."}
    KossonSILVER0015 = Kanon {"Kanon  5.  15", "Der JHWH und der Prophet sind Mensch also, und sie belehren in eigener Pflichterkennung und Pflichterfüllung die Gesetze und Gebote der Schöpfung."}
    KossonSILVER0016 = Kanon {"Kanon  5.  16", "Der JHWH und der Prophet belehren in Weisheit, Liebe, Wahrheit und Wissen die Gesetze und Gebote der Schöpfung, ohne Stellvertreter der Schöpfung zu sein."}
    KossonSILVER0017 = Kanon {"Kanon  5.  17", "Die Schöpfung ist unendliche Kraft, unmessbare Grösse und unmessbares Geheimnis, und keine menschliche Kraft vermag sich ihr zu nähern."}
    KossonSILVER0018 = Kanon {"Kanon  5.  18", "Kommunikation vermag der Geist im Menschen mit der Schöpfung zu betreiben, denn er allein ist ein energiegleiches Teilstück der Schöpfung selbst."}
    KossonSILVER0019 = Kanon {"Kanon  5.  19", "Die Schöpfung allein soll verehret (geehrwürdigt) und angebetet (geehrt) werden, und ihr allein sei Preis (Ehre) und Dank."}
    KossonSILVER0020 = Kanon {"Kanon  5.  20", "Der JHWH und der Prophet sollen respektieret und ehrfürchtig geachtet, jedoch nicht verehret （geehrwürdigt) und nicht angebetet werden."}
    KossonSILVER0021 = Kanon {"Kanon  5.  21", "Anbetung(Achtung)und Verehrung (Erwürdigung)gebühret allein der Schöpfung,und sie sei gepriesen in Dankbarkeit."}

    //Kanon 6
    KossonSILVER0022 = Kanon {"Kanon  6.   1", "Im Namen der Schöpfung, die da ist die alleinige Schöpfung in ihrem Himmel (Universum), und die da ist das Leben und das SEIN (Schöpfungsleben)."}
    KossonSILVER0023 = Kanon {"Kanon  6.   2", "So sprichet der Weisheitskönig der irdischen Menschengeschlechter zum Propheten, der da ist der wahrliche Prophet und der als Mittler und Künder das Buch der Wahrheit als Buch des wahrlichen Wortes und der Lehre empfanget und den Menschengeschlechtern der Erde bringet, denen es einheitlich als Wegeleitung und Ratgebung und Befolgungsrichtlinie gelte:"}
    KossonSILVER0024 = Kanon {"Kanon  6.   3", "Es besagen die Gesetze und Gebote der Schöpfung, dass eine jegliche Herde eines Leittieres bedarf."}
    KossonSILVER0025 = Kanon {"Kanon  6.   4", "Und ein jegliches Heer bedarf eines Führers."}
    KossonSILVER0026 = Kanon {"Kanon  6.   5", "Und ein jegliches Volk bedarf einer weisen Leitung."}
    KossonSILVER0027 = Kanon {"Kanon  6.   6", "Und ein jegliches Menschengeschlecht bedarf eines weisen Königs."}
    KossonSILVER0028 = Kanon {"Kanon  6.   7", "Ein jegliches Menschengeschlecht bedarf eines weisen, wissenden, in Liebe, Pflicht, Können und in Erkenntnis und Ratgebung gebildeten Königs, so also des Weisheitskönigs."}
    KossonSILVER0029 = Kanon {"Kanon  6.   8", "Und es sind genennet der JHWH und der Prophet als Weisheitskönige, über denen allein die Schöpfung aller Dinge als unmessbares Geheimnis waltet."}
    KossonSILVER0030 = Kanon {"Kanon  6.   9", "Siehe also: Der JHWH und der Prophet wachen über die Menschengeschlechter, die sie leiten und führen durch jene, die als Künder und Führer ihre Wünsche, ihre Worte und ihr Buch der Gesetze und Gebote der Schöpfung zu den Menschen der Erden bringen und diese belehren."}
    KossonSILVER0031 = Kanon {"Kanon  6.  10", "Die Mittler, Künder und Führer der Wahrheit und der Gesetze und Gebote der Schöpfung sind der JHWH und der Prophet und die Geistführer des Propheten, die da die Belehrung des unwissenden Menschen als Pflicht erfüllen und die da selbst die Kraft grösster Weisheit, der Wahrheit, des Wissens und der Liebe tragen."}
    KossonSILVER0032 = Kanon {"Kanon  6.  11", "Der JHWH ist nicht Stellvertreter der Schöpfung, und es ist nicht stellvertretend der Prophet für den JHWH."}
    KossonSILVER0033 = Kanon {"Kanon  6.  12", "Es kündet der Prophet für die Menschengeschlechter einer Erde seine Lehre und die Lehre des JHWH zu je seiner Zeit, und kein anderer Prophet und kein anderer JHWH ist gleichsam neben ihnen zu ihrer Zeit."}
    KossonSILVER0034 = Kanon {"Kanon  6.  13", "Und es künden der JHWH und der Prophet die Lehre des Geistes, die da ist die Lehre der Schöpfung aus ihren Gesetzen und Geboten, aufgetragen dem JHWH und dem Propheten durch eigene Pflichterkennung zur Erfüllung."}
    KossonSILVER0035 = Kanon {"Kanon  6.  14", "Und also ordnet der Prophet sich Geistführer ein, die bei ihm lernend und danach unter den Menschen belehrend sind, so sie als wissende und pflichterfüllende Geistführer die Massen der Menschengeschlechter belehrend unterrichten und führen."}

    //Kanon 7
    KossonSILVER0036 = Kanon {"Kanon  7.   1", "Im Namen der Schöpfung, die da ist gerecht und grossmütig."}
    KossonSILVER0037 = Kanon {"Kanon  7.   2", "Preis (Ehre), Verehrung (Ehrwürdigung)und Anbetung (Achtung) sei der Schöpfung allein."}
    KossonSILVER0038 = Kanon {"Kanon  7.   3", "Die Schöpfung ist unsichtbar, ungreifbar und unfühlbar."}
    KossonSILVER0039 = Kanon {"Kanon  7.   4", "Es ist die Schöpfung aber empfindbar."}
    KossonSILVER0040 = Kanon {"Kanon  7.   5", "Allein wahrliches Wissen, wahrliche Weisheit, wahrliche Wahrheit und wahrliche Liebe machen den Menschen verstehend, erfassend und erkennend in der Beweisung der Existenz der Schöpfung."}
    KossonSILVER0041 = Kanon {"Kanon  7.   6", "Also bleibet die Existenz jenem Menschen unbeweisbar, der die Beweiskraft nicht findet in wahrlicher Wahrheit, in wahrlicher Liebe, in wahrlicher Weisheit und in wahrlichem Wissen."}
    KossonSILVER0042 = Kanon {"Kanon  7.   7", "Der JHWH ist sichtbar und erfassbar, denn er ist Mensch, er aber hütet sein Sichtbares und Erfassbares, so er sich dem Blicke und dem Griffe jener Menschen enthaltet, die noch der Belehrung der Gesetze und Gebote der Schöpfung bedürfen."}
    KossonSILVER0043 = Kanon {"Kanon  7.   8", "Also bleibet der JHWH den noch zu Belehrenden unsichtbar und ungreifbar, und nur die Propheten erlangen den Anblick in die Güte seines Angesichts."}
    KossonSILVER0044 = Kanon {"Kanon  7.   9", "Unwissende und noch zu belehrende Menschen verehren in irrem Streben Höheres, solches, das über ihnen stehet und wissender und könnender ist; also verehret der Unwissende auch Menschen, die wissender und wahrlich weise sind."}
    KossonSILVER0045 = Kanon {"Kanon  7.  10", "Und der JHWH ist Mensch, weise und wissend, und er darf nicht verehret (bewundert) werden, so bleibet er dem noch zu belehrenden Menschen unsichtbar und ungreifbar, damit er nicht verehret und nicht angebetet werde."}
    KossonSILVER0046 = Kanon {"Kanon  7.  11", "Und wahrlich, Verehrung (Bewunderung)und Anbetung (Achtung) gebühret allein der Schöpfung."}
    KossonSILVER0047 = Kanon {"Kanon  7.  12", "Der JHWH ist Mensch, und als solcher wünschet und duldet er keine Verehrung(Ehrwürdigung) und Anbetung (keinen Ruhm), so aber auch nicht die Propheten, die sich dagegen zu erwehren haben und sich so nicht direkt im Volke bewegen, denn sie sind doch bei den Menschengeschlechtern tätig, die in irrem Streben Verehrung und Anbetung suchen."}
    KossonSILVER0048 = Kanon {"Kanon  7.  13", "Es haben sich zu erwehren die Propheten gegen Verehrung und Anbetung, denn auch über ihnen schwebet als unmessbares Geheimnis die Schöpfung, und ihr allein gebühret Verehrung und Anbetung."}
    KossonSILVER0049 = Kanon {"Kanon  7.  14", "Also bleibet der JHWH dem noch zu belehrenden Menschen unsichtbar und ungreifbar, und so auch alle seine ihm Angeordneten, die Helfer und Wächter, so die noch zu Belehrenden nicht zur Verehrung und Anbetung neigend werden und so der Mensch allein durch Verstand und Vernunft und durch Wissen, Wahrheit, Liebe und Weisheit um seine Existenz wissend werde, gleichermassen er auch zu tun hat im Wissen um die Existenz der Schöpfung."}

    //Kanon 8
    KossonSILVER0050 = Kanon {"Kanon  8.   1", "Im Namen der Schöpfung,die da kreieret hat die Heere allen Lebens."}
    KossonSILVER0051 = Kanon {"Kanon  8.   2", "Preis (Ehre) sei der Schöpfung, die da ist die Schöpfung."}
    KossonSILVER0052 = Kanon {"Kanon  8.   3", "Und es sprichet der JHWH und der Prophet, die da wissen um den Anfang und um die Gesetze und Gebote:"}
    KossonSILVER0053 = Kanon {"Kanon  8.   4", "«Am Anfang war die Eiform der Schöpfungsspirale, hervorgebracht durch die ldee der Urschöpfung."}
    KossonSILVER0054 = Kanon {"Kanon  8.   5", "Die Schöpfungsspirale war geistige Urform, winzig klein und nur von der Grösse eines Flohs."}
    KossonSILVER0055 = Kanon {"Kanon  8.   6", "In der Spirale aber rotierete und pulsierete die Geistenergieform der werdenden Schöpfung, gewaltig und ausdehnungsgewillet."}
    KossonSILVER0056 = Kanon {"Kanon  8.   7", "Die Energieform wuchs und wurde zur geballeten Kraft, und explodierete in grellstem Blitzefeuer."}
    KossonSILVER0057 = Kanon {"Kanon  8.   8", "Ein Sohar zwischen den zahllosen Universen des Absolutumraumes."}
    KossonSILVER0058 = Kanon {"Kanon  8.   9", "Mit gewaltiger Ausdehnungskraft trieb die Schöpfungsspirale auseinander, sich ausweitend mit vielfacher Geschwindigkeit des Lichtes, stetig wachsend und schöpferisch-universellen Raum schaffend."}
    KossonSILVER0059 = Kanon {"Kanon  8.  10", "Die Schöpfung ward geboren, erdacht durch die Idee der Urschöpfung, und kreieret durch die eigene Kraft."}
    KossonSILVER0060 = Kanon {"Kanon  8.  11", "Aus der endlosen Dauer in der Urzelle befreiet, lösete sich in ihr der Griff der Fesseln - und die Schöpfung begann zu denken in bewusster Form."}
    KossonSILVER0061 = Kanon {"Kanon  8.  12", "Es entwickelten sich Vernunft und Verstand und die Empfindung und das Gemüt."}
    KossonSILVER0062 = Kanon {"Kanon  8.  13", "So schuf die Idee den Gedanken, die da ist vorgegeben durch das Urewige, durch das Absolute Absolutum, welches der Ursprung aller Dinge ist."}
    KossonSILVER0063 = Kanon {"Kanon  8.  14", "Das Ur ist das sich selbst Kreierete aus dem absoluten Nichts der endlosen Dauer, die da genennet ist als Zeitlosigkeit."}
    KossonSILVER0064 = Kanon {"Kanon  8.  15", "Die Idee des Urewigen aber ist gegeben im evolutiven Sinn des Daseins, im Werden und Vergehen, bis hin zum endgültigen Leben des allzeitlichen und zeitlosen SEIN (Schöpfungsleben) im Absoluten Absolutum."}
    KossonSILVER0065 = Kanon {"Kanon  8.  16", "Die Schöpfung lebete, dachte, hatte Gemüt, Empfindung, Verstand, Vernunft und die ldee des Lebens."}
    KossonSILVER0066 = Kanon {"Kanon  8.  17", "Und sie ward kraftvoll und machtvoll in wahrlicher Liebe, um sich auszudehnen und zu weiten."}
    KossonSILVER0067 = Kanon {"Kanon  8.  18", "Mit ihrer Kraft wandelte sie ihre Gedanken und ihre Idee zur Werdung, so sie daraus schuf Raum und Zeit."}
    KossonSILVER0068 = Kanon {"Kanon  8.  19", "Sie schuf das Universum allso, das in seiner Weite leer und nur angefüllet war mit der geisti gen Kraftenergie der sich ausdehnenden Schöpfung."}
    KossonSILVER0069 = Kanon {"Kanon  8.  20", "Also schuf die Schöpfung am Anfang die leere Weite des Himmels (Universums), in dem noch endlose Dauer herrschete und noch kein Raum gegenwärtig geworden war."}
    KossonSILVER0070 = Kanon {"Kanon  8.  21", "Doch die Idee der Wandlung zum Werden war kreieret, und es wurde die endlose Dauer zur Zeit und die leere Weite zum Raum."}
    KossonSILVER0071 = Kanon {"Kanon  8.  22", "Doch der Raum ward leer und die Zeit stand still also."}
    KossonSILVER0072 = Kanon {"Kanon  8.  23", "Die Idee aber brachte Bewegung hervor."}
    KossonSILVER0073 = Kanon {"Kanon  8.  24", "Es spaltete sich die Dauer in winzigste Impulseinheiten, die fortan den Raum in Wellen durcheileten."}
    KossonSILVER0074 = Kanon {"Kanon  8.  25", "Die sich bewegende Zeit ward geboren, und lebend schuf sie Vergangenheit und Zukunft."}
    KossonSILVER0075 = Kanon {"Kanon  8.  26", "Die eilende Zeit setzete die im Raum schwebenden geistigen Schöpfungsenergien in Bewegung, und Rotation entstand."}
    KossonSILVER0076 = Kanon {"Kanon  8.  27", "Bewegend und rotierend pulsiereten die geistigen Energien, und flockige Materie kreierete daraus."}
    KossonSILVER0077 = Kanon {"Kanon  8.  28", "Also ward getan der erste Schritt der Materiebildung."}
    KossonSILVER0078 = Kanon {"Kanon  8.  29", "Weiter entwickelte sich die ldee der Schöpfung, die urewig ihr eingegeben ward durch die Urschöpfung, der wiederum die Idee ursprünglich ward durch das Absolute Absolutum."}
    KossonSILVER0079 = Kanon {"Kanon  8.  30", "Noch war im Himmel (Universum) nur eilende Zeit, Bewegung, Rotation, Pulsierung und flokkige Materie."}
    KossonSILVER0080 = Kanon {"Kanon  8.  31", "Der Raum war wirr und wüst und hing in Finsternis, und keine Feste war vorhanden."}
    KossonSILVER0081 = Kanon {"Kanon  8.  32", "Also weitete sich die ldee der Schöpfung,und sie schuf die Rotation zur Verdichtung."}
    KossonSILVER0082 = Kanon {"Kanon  8.  33", "Die flockige Materie rotierete und pulsierete zu Spiralen, der Eiform der Schöpfungsspirale gleich, und verdichtete sich, wodurch Wärme entstand, die andersartig war als die Kälte des Raumes."}
    KossonSILVER0083 = Kanon {"Kanon  8.  34", "Die sich verdichtende Flockenmaterie der Geistenergie wurde zur Galaxie, so auch sie ward geboren, glühend und strahlend hell, wodurch die Finsternis des Raumes gebrochen wurde."}
    KossonSILVER0084 = Kanon {"Kanon  8.  35", "Geboren also ward die Urgalaxie."}
    KossonSILVER0085 = Kanon {"Kanon  8.  36", "Aus der Urgalaxie entstand neues Werden, neue verdichtete Materie, die in Bewegung war, rotierete und pulsierete."}
    KossonSILVER0086 = Kanon {"Kanon  8.  37", "Die Materie grober Form entstand, in Galaxien neuer Art, die vielfältig und zahllos den weiten Raum belebeten."}
    KossonSILVER0087 = Kanon {"Kanon  8.  38", "Noch war dadurch aber die ldee der Schöpfung nicht vollendet, denn noch ward kein Leben zur Evolution erschaffen, und nichts, das da kreuchet und fleuchet."}
    KossonSILVER0088 = Kanon {"Kanon  8.  39", "Und noch waren keine Gesetzmässigkeiten und Gebotsmässigkeiten, in die das Leben und die Existenz aller Dinge eingeordnet sein sollten, nicht für das Dasein und nicht für die Evolution."}
    KossonSILVER0089 = Kanon {"Kanon  8.  40", "Also dachte die Schöpfung und kreierete allumfassende Gesetze und Gebote, in die alles Leben, alle Evolution, das Werden und das Vergehen und das SEIN (Schöpfungsleben) eingeordnet sein sollten."}
    KossonSILVER0090 = Kanon {"Kanon  8.  41", "Es ward damit erschaffen die Siebenheit der schöpferischen Gesetzmässigkeiten und Gebotsmässigkeiten."}
    KossonSILVER0091 = Kanon {"Kanon  8.  42", "Also waren damit gegeben die Richtlinien allen Lebens und aller Evolution, und es waren gegeben die Gesetze und Gebote für alles Bestehende, Werdende, Vergehende und für das SEIN（Schöpfungsleben)."}
    KossonSILVER0092 = Kanon {"Kanon  8.  43", "Und die Schöpfung ordnete ihre Gesetzmässigkeiten und Gebotsmässigkeiten ein in alles bereits von ihr Erschaffene und in alles, was von ihr noch erschaffen würde."}
    KossonSILVER0093 = Kanon {"Kanon  8.  44", "Sie verankerte ihre Gesetze und Gebote in allem Bestehenden und Werdenden, auf dass allgrosszeitlich alles Heer des Lebens und des Existenten dieser Ordnung eingebettet sei."}
    KossonSILVER0094 = Kanon {"Kanon  8.  45", "Es nahmen die Gesetzmässigkeiten und Gebotsmässigkeiten der allumfassenden siebenheitlichen Ordnung ihren Lauf, und neuerlich entstand daraus auch andersartiges Leben."}
    KossonSILVER0095 = Kanon {"Kanon  8.  46", "Die siebenheitliche Ordnung ward geboren, die da besaget, dass alles Leben, alles Bestehen, alles Werden und Vergehen, alle Evolution und das SEIN (Schöpfungsleben) sich je in sieben aufeinanderfolgenden Stufen und Ebenen ergänzen würde."}
    KossonSILVER0096 = Kanon {"Kanon  8.  47", "Also besaget die siebenheitliche Ordnung der Schöpfungsgesetze und Schöpfungsgebote, dass alles einem Lauf und Fortgang eingeordnet sei, der von seinem Anbeginn bis zum Ziele sieben Stufen umfasset oder sieben Ebenen also."}
    KossonSILVER0097 = Kanon {"Kanon  8.  48", "Noch nicht ist damit aber die siebenheitliche Ordnung erfüllet, denn eine jegliche Ebene wiederum ist unterteilet in weitere sieben Ebenen, und diese wiederum in sieben niedrigere Ebenen also, und dies in siebenmaliger aufeinanderfolgender Form."}
    KossonSILVER0098 = Kanon {"Kanon  8.  49", "Es ergibet sich so die siebenheitliche Ordnung der 7 x 7fachen Entwicklung, die da besaget, dass allem Heer allen Lebens sieben direkt schöpfungsbezogene Entwicklungsebenen gegeben sind."}
    KossonSILVER0099 = Kanon {"Kanon  8.  50", "Und eine jegliche schöpfungsbezogene Entwicklungsebene ist unterteilet in sieben niedrigere Entwicklungsebenen."}
    KossonSILVER0100 = Kanon {"Kanon  8.  51", "Und eine jegliche niedrigere Entwicklungsebene ist unterteilet in sieben noch niedrigere Entwicklungsebenen."}
    KossonSILVER0101 = Kanon {"Kanon  8.  52", "Und eine jegliche niedrigere Entwicklungsebene ist unterteilet in sieben noch tiefere Entwicklungsebenen, das so fortlaufend ist bis in das siebente Glied."}
    KossonSILVER0102 = Kanon {"Kanon  8.  53", "Also ergibet sich die Zahl aller Ebenen in 7 x 7 x 7 x 7 x 7 x 7 x 7, die alle Dinge und alles Bestehen durchweben, jegliches nach seiner Art."}
    KossonSILVER0103 = Kanon {"Kanon  8.  54", "Und es entstand aus der siebenheitlichen Ordnung das neue und andersartige Leben, das da ward dichtere und grobstoffliche Materie."}
    KossonSILVER0104 = Kanon {"Kanon  8.  55", "Die glühenden, rotierenden und pulsierenden Spiralarme der vielartigen neuen Galaxien schufen in Verdichtung grosse selbstrotierende Ballungen, die sich verdichteten zu grossen Gebilden und die sich vom Zentrum der Galaxien entfernten."}
    KossonSILVER0105 = Kanon {"Kanon  8.  56", "Aus den grossen, rotierenden, pulsierenden und glühenden Gebilden waren alsbald geworden Verdichtungen, die sich bildeten zu Sonnen und Planeten und die hinausgeschleudert waren in die Weiten des Raumes."}
    KossonSILVER0106 = Kanon {"Kanon  8.  57", "Und die Sonnen und Planeten waren da vereinzelt im Raume, und sie waren vereinet und sich gegenseitig umkreisend in Gruppen und Systemen also."}
    KossonSILVER0107 = Kanon {"Kanon  8.  58", "Also belebeten sie den Raum des Himmels (Universums) in unterschiedlicher Grösse und in vielfältiger Form und Art, und in verschiedensten Farben, so auch der Raum des Himmels (Universums) vielfältig ward an Farben also."}
    KossonSILVER0108 = Kanon {"Kanon  8.  59", "Und die Sonnen erstrahleten in eigenem Feuer und spendeten Licht und Wärme an die sie umkreisenden Planeten, die selbst zu Teilen und oberflächlich erkalteten und vereinzelt fähig wurden, eigenes Leben zu entwickeln, wie dies die siebenheitliche Ordnung der Schöpfung bestimmete."}
    KossonSILVER0109 = Kanon {"Kanon  8.  60", "Also ward es gegeben durch die Gesetze und Gebote der Schöpfung, dass Planeten Leben entwickelten und Leben tragen sollten, jeglicher Planet nach seiner Art."}
    KossonSILVER0110 = Kanon {"Kanon  8.  61", "So kreierete das Leben in mannigfacher Form und Art auf vereinzelten Welten, die da genannt sind Erden."}
    KossonSILVER0111 = Kanon {"Kanon  8.  62", "Und es gebar das erste Leben als Flora, die die Festen der Erden bewuchsen."}
    KossonSILVER0112 = Kanon {"Kanon  8.  63", "Und alsbald gebar auch das Leben der Fauna, die da ward die Welt des Tierreiches, nachdem die Flora gediehen war, die da ist das Pflanzenreich."}
    KossonSILVER0113 = Kanon {"Kanon  8.  64", "Also ward erst kreieret die Flora, und es folgete ihr dann die Fauna, gemäss den siebenheitlichen Ordnungsgesetzen der Schöpfung."}
    KossonSILVER0114 = Kanon {"Kanon  8.  65", "Und es gebar als letzte Lebensform der Mensch, als bereits lebendig waren die Flora und die Fauna, und damit gegeben ward die Möglichkeit des Lebens des Menschen, der da ward die höchstentwickelte und die selbständig denkende Lebensform, vorbedacht durch die Idee der Schöpfung."}
    KossonSILVER0115 = Kanon {"Kanon  8.  66", "Am Anfang ward die Schöpfung, die da schuf den Himmel (Universum), den Raum und die Zeit und die Festen darinnen."}
    KossonSILVER0116 = Kanon {"Kanon  8.  67", "Auf den Festen aber kreierete mannigfaches Leben, so der Idee des Werdens zum SEIN（Schöpfungsleben) im Absoluten Absolutum Genüge getan ward und der Weg der Evolution eröffnet sei."}
    KossonSILVER0117 = Kanon {"Kanon  8.  68", "Und so sei der Gang aller Dinge, wenn alle Heere des Lebens die durch die Schöpfung gegebenen Gesetze und Gebote befolgen."}
    KossonSILVER0118 = Kanon {"Kanon  8.  69", "Die Gesetze und Gebote, die da sind urzeitlich gegeben und die da sind dem Wissenden und Weisen in wahrlicher Liebe allüberall erkennbar in allgrosszeitlich gleichbleibender und unveränderlich dauernder Form."}
    KossonSILVER0119 = Kanon {"Kanon  8.  70", "Also ward auch die Erde geschaffen nach den Gesetzen und Geboten der Schöpfung, so aber ward auch erschaffen die Sonne und die anderen Planeten also."}
    KossonSILVER0120 = Kanon {"Kanon  8.  71", "Am Anfang ward die Sonne des neuen Systems, und sie spendete Wärme und Licht."}
    KossonSILVER0121 = Kanon {"Kanon  8.  72", "Und weit ausserhalb der Sonne verdichteten sich Materiengebilde, und es entstanden die Planeten."}
    KossonSILVER0122 = Kanon {"Kanon  8.  73", "Das Licht der Sonne erhellete die Finsternis und die rotierenden Planeten, so ward Licht auf ihnen also, und Dunkelheit auf der sonneabgewandten Seite."}
    KossonSILVER0123 = Kanon {"Kanon  8.  74", "Also ward auf den Planeten Tag und Nacht, und auf der Erde also."}
    KossonSILVER0124 = Kanon {"Kanon  8.  75", "Und gemäss den siebenheitlichen Gesetzen und Geboten der Schöpfung kreierete auf der Erde mannigfaches Leben."}
    KossonSILVER0125 = Kanon {"Kanon  8.  76", "Es ward geboren Flora und Fauna, ein jegliches Leben nach seiner Art."}
    KossonSILVER0126 = Kanon {"Kanon  8.  77", "Und das Leben auf der Erde ward gegeben, weil die Welt zu Teilen aus Feste und zu Teilen aus Wasser ward, aus denen sich erstlich und letztlich alles Leben am Anfang kreierete."}
    KossonSILVER0127 = Kanon {"Kanon  8.  78", "Und das Leben kreierete bis hin zum Menschen, der sich entwickelte als separate und eigenlinientreue und höchstentwickelte schöpferische Lebensform, heraus aus Ursubstanzen der Erde und aus deren Leben in substanzenverbindender Form."}
    KossonSILVER0128 = Kanon {"Kanon  8.  79", "Und es kreierete das menschliche Leben in substanzenverbindender Form der Erde und des bestehenden Lebens der Flora und Fauna, ohne direkte Gattungsgemeinschaft mit florischen oder faunaischen Lebensformen."}
    KossonSILVER0129 = Kanon {"Kanon  8.  80", "Also ward geboren die Gattung Mensch als eigene und spezielle Art, in keinerlei direkter Gemeinschaft mit florischem oder faunaischem Leben also."}
    KossonSILVER0130 = Kanon {"Kanon  8.  81", "Und die Lebensform Mensch auf der Erde entwickelte Vernunft und Verstand, das Fühlen und das Denken und das Sprechen also."}
    KossonSILVER0131 = Kanon {"Kanon  8.  82", "Es bildete der Mensch Worte und reihete sie aneinander zu bestimmten Begriffsfolgerichtigkeiten, und es entstand daraus der Sinn der aufeinanderfolgenden Worte und deren Verbindungen."}
    KossonSILVER0132 = Kanon {"Kanon  8.  83", "Es entstand der Wert der Sprache, so der Mensch sich durch Worte gegenseitig verstehend machen konnte."}
    KossonSILVER0133 = Kanon {"Kanon  8.  84", "Es ward kreieret die Sprache also."}
    KossonSILVER0134 = Kanon {"Kanon  8.  85", "Und so der Mensch der Sprache mächtig geworden war, nennete er den Raum über der Welt Himmel und die trockene Feste Erde, und das feuchte Nass nennete er Wasser."}
    KossonSILVER0135 = Kanon {"Kanon  8.  86", "Es ward also beschlossen durch die Gesetze und Gebote der Schöpfung, dass auf bestimmten Gestirnen Leben mannigfacher Art und Gattung entstehe, so die Planeten in ihrer Art lebentragungsfähig waren."}
    KossonSILVER0136 = Kanon {"Kanon  8.  87", "Und also war der Planet Erde einer von ihnen."}
    KossonSILVER0137 = Kanon {"Kanon  8.  88", "Die Gesetze und Gebote der Schöpfung besageten, dass auch auf der Erde aufgehen sollten Heere des Lebens, jegliches nach seiner Art."}
    KossonSILVER0138 = Kanon {"Kanon  8.  89", "Und die Gesetze und Gebote der Schöpfung besageten, dass auf der Erde sollten aufgehen Gras und allerhand Kraut, jegliches nach seiner Art."}
    KossonSILVER0139 = Kanon {"Kanon  8.  90", "Und es sollten aufgehen vielartige Bäume, Blumen, Pflanzen und Sträucher, und Samen sollten sie tragen und Früchte in grossen Mengen, und duftend sollten sie sein, schmeckend und wohlriechend, jegliche nach ihrer Art."}
    KossonSILVER0140 = Kanon {"Kanon  8.  91", "Und aus den Wassern und dem Erdreich entwickelte sich alles, wie es ward vorgesehen und erdenket durch die Idee der Schöpfung."}
    KossonSILVER0141 = Kanon {"Kanon  8.  92", "Und die Erde liess aufgehen Gras und allerlei Kraut, Blumen, Pflanzen, Bäume und Sträucher und allerlei Getier."}
    KossonSILVER0142 = Kanon {"Kanon  8.  93", "Und alle entstandenen Heere des Lebens trugen eigene Samen und Früchte, so sie sich in weiterer Folge selbst besameten und fortpflanzeten, jegliches nach seiner Art."}
    KossonSILVER0143 = Kanon {"Kanon  8.  94", "Und es ward Atemraum über der Feste der Erde und unter dem Himmel, so atmen konnte jegliches Leben, und jegliches nach seiner Art."}
    KossonSILVER0144 = Kanon {"Kanon  8.  95", "Und es waren Lichter gesetzet in die Dunkelheit des Himmels über der Erde, so des Tags die Sonne Licht und Wärme gab und des Nachts ferne Gestirne bescheidenes Licht verbreiteten."}
    KossonSILVER0145 = Kanon {"Kanon  8.  96", "So ward gegeben Tag und Nacht, entstehend durch die Zeichen des Lichts und der Dunkelheit."}
    KossonSILVER0146 = Kanon {"Kanon  8.  97", "Die Erdenfeste kreisete im Allraum um die Sonne und um sich selbst, und also ward dadurch gegeben der Wandel von Tag und Nacht."}
    KossonSILVER0147 = Kanon {"Kanon  8.  98", "Und es waren gegeben durch die Drehung die Zeiten der Tage und die Stunden und die Zeiten der Jahre, so aber auch die Wandlung des Werdens und Vergehens des gesamten Lebens."}
    KossonSILVER0148 = Kanon {"Kanon  8.  99", "Und also ward damit auch gegeben der Frühling, der Sommer, der Herbst und der Winter, die da sind in der Wandlungsform die Geburt, das tätige Leben, der besinnliche Lebensabend und der Tod, der wie ein Tiefschlaf zum neuerlichen Wiederleben führt."}
    KossonSILVER0149 = Kanon {"Kanon  8. 100", "Also herrschete auf der Erde Werden und Vergehen, gegeben durch die siebenheitlichen Gesetze und Gebote der Schöpfung."}
    KossonSILVER0150 = Kanon {"Kanon  8. 101", "Es ward Werden und Vergehen und stetig wieder neues Werden und Vergehen"}
    KossonSILVER0151 = Kanon {"Kanon  8. 102", "Und es ward Wasser und Feste, Feste und Himmel,Tag und Nacht, Jahr um Jahr, und es ward Jahreszeitenwandlung, vielfältiges Leben, Besamung und Geburt und Vergehen und Wiedergeburt in endloser Folge, und jegliches nach seiner Art."}
    KossonSILVER0152 = Kanon {"Kanon  8. 103", "Also ward existent und lebend geworden die Erde, die Wasser, die Himmel (Universum, Ebenen und Dimensionen) und die Gestirne, und jegliche Heere des Lebens und nach ihrer Art also."}
    KossonSILVER0153 = Kanon {"Kanon  8. 104", "Und jegliches Leben ward erfüllet und belebet von winzigen Teilen schöpferischer Kraft und schöpferischen Geistlebens, das sich entwickeln soll zur möglichsten Vollkommenheit, wodurch erfüllet sein sollen die Gesetze und Gebote der Schöpfung, bis hin in alle Allgrosszeit."}
    KossonSILVER0154 = Kanon {"Kanon  8. 105", "Also ward vollendet und dem Leben und der Evolution übergeben die Erde und die Gestirne und alle Heere des Lebens, derweil der Himmel (Universum) weiter sich ausweitet und erst dann innehalten wird, wenn die Zeit der Zeit gekommen ist."}
    KossonSILVER0155 = Kanon {"Kanon  8. 106", "Also ward vollendet die Erde, die Himmel (Universum, Ebenen und Dimensionen),die Gestirne und alle Heere des Lebens, das Licht und die Dunkelheit, die Kälte und die Wärme und die Wasser und alles übrige, das da ist existent."}
    KossonSILVER0156 = Kanon {"Kanon  8. 107", "Also ist alles geworden, kreieret und geboren und erschaffen am Anfang, von Anbeginn der Idee der Urschöpfung zur Selbstkreation der Schöpfung."}
    KossonSILVER0157 = Kanon {"Kanon  8. 108", "Aus der Schöpfung aber entsprang und gebar der Himmel (Universum), die Galaxien, die Gestirne und alle Heere des Lebens, die da kreuchen und fleuchen, in Erfüllung der Gesetze und Gebote der Schöpfung, die da sind von Gültigkeit von Allgrosszeit zu Allgrosszeit und ewiglich."}
    KossonSILVER0158 = Kanon {"Kanon  8. 109", "Also ward geboren die Schöpfung selbst durch die Idee der Urschöpfung also."}
    KossonSILVER0159 = Kanon {"Kanon  8. 110", "Und also ward geboren der Himmel (Universum) durch die Idee der Schöpfung."}
    KossonSILVER0160 = Kanon {"Kanon  8. 111", "Und also waren geboren die Gesetzmässigkeiten und Gebotsmässigkeiten der Siebenheit durch die ldee der Schöpfung."}
    KossonSILVER0161 = Kanon {"Kanon  8. 112", "Und also waren geboren die Festen, die Galaxien und Gestirne im Raum des Himmels (Universums), und so auch Raum und Zeit, Licht und Dunkelheit durch die Idee der Schöpfung."}
    KossonSILVER0162 = Kanon {"Kanon  8. 113", "Und endlich ward geboren die vielfältige und mannigfache Heerschaft des Lebens, die da evolutionierend das Ziel des SEINs (Schöpfungslebens) erstrebt."}
    KossonSILVER0163 = Kanon {"Kanon  8. 114", "Es geschah so im Willen und nach der Idee des Absoluten Absolutums."}
    KossonSILVER0164 = Kanon {"Kanon  8. 115", "Und es geschah so im Willen und nach der ldee der Urschöpfung."}
    KossonSILVER0165 = Kanon {"Kanon  8. 116", "Und es geschah so im Willen und nach der ldee der Schöpfung."}
    KossonSILVER0166 = Kanon {"Kanon  8. 117", "Damit das Leben in zielstrebiger Evolution gedeihe."}
    KossonSILVER0167 = Kanon {"Kanon  8. 118", "Und damit das Leben durch Evolution im SEIN (Schöpfungsleben) die relative Vollkommenheit finde."}
    KossonSILVER0168 = Kanon {"Kanon  8. 119", "Im Namen der Schöpfung, die da ist Wahrheit, und die da ist Weisheit, und die da ist wahrheitliche Liebe, dies ist die Wahrheit vom Anbeginn des Anbeginns.»"}

    //Kanon 9
    KossonSILVER0169 = Kanon {"Kanon  9.   1", "Im Namen der Schöpfung, der Weisen, der Gerechten."}
    KossonSILVER0170 = Kanon {"Kanon  9.   2", "Preis (Ehre) sei der Schöpfung, die da Himmel und Erde erschaffen und Dunkelheit und Licht geboren hat."}
    KossonSILVER0171 = Kanon {"Kanon  9.   3", "Preis (Ehre) sei der Schöpfung, denn sie ist es, die da den Menschen erdacht und kreieret und ihm ein evolutives Ziel gesetzet hat."}
    KossonSILVER0172 = Kanon {"Kanon  9.   4", "Und sie ist die Schöpfung, die da existent ist in allem, in den Himmeln, in den Gestirnen und Galaxien, in den Erden und in allem Heer des Lebens, das da ist gegeben jegliches nach seiner Art."}
    KossonSILVER0173 = Kanon {"Kanon  9.   5", "Und sie ist erkennbar in allem und in jeglichem, innerlich und äusserlich."}
    KossonSILVER0174 = Kanon {"Kanon  9.   6", "Und sie hat das Lebensziel und das Endziel des SEINs (Schöpfungslebens) bestimmet, für alles, was da kreuchet und fleuchet, und für alles, was lebet und existieret."}
    KossonSILVER0175 = Kanon {"Kanon  9.   7", "Der Erdenmensch aber zweifelt daran, und es setzen jene sich über sie hinweg, die unwissend sind und selbstherrlich und dumm."}
    KossonSILVER0176 = Kanon {"Kanon  9.   8", "Sie kennen nicht Ehrfurcht (Ehrung, Ehrerbietung, Ehrwürdigung, Ehrwürdigkeit) und nicht Respekt, nicht die Wahrheit, nicht die Liebe und weder Wissen noch Weisheit."}
    KossonSILVER0177 = Kanon {"Kanon  9.   9", "Sie sind Gefangene ihrer selbst und voller Ungerechtigkeit, Übellaunigkeit und Selbstsucht und Selbstmitleid."}
    KossonSILVER0178 = Kanon {"Kanon  9.  10", "Und sie sind gefangen in ihrem Egoismus und in ihrem irren und falschen Denken."}
    KossonSILVER0179 = Kanon {"Kanon  9.  11", "Sie setzen der Schöpfung anderes gleich und treiben damit Verehrung und Anbetung mit Götzen und Göttern von toter und lebendiger Art."}
    KossonSILVER0180 = Kanon {"Kanon  9.  12", "Und sie lügen ihnen Schöpfungskraft zu, und sie fürchten sich vor ihnen."}
    KossonSILVER0181 = Kanon {"Kanon  9.  13", "Und sie setzen Menschen über die Schöpfung, die sie als Gewaltherrscher verehren und anbeten (anbetteln und anflehen) und die da sind genennet beim Erdenmenschen Gott."}
    KossonSILVER0182 = Kanon {"Kanon  9.  14", "Und sie sprechen irre ihren Göttern und Götzen Macht zu, und sie setzen sie der Schöpfung gleich."}
    KossonSILVER0183 = Kanon {"Kanon  9.  15", "Die Schöpfung ist aber die Schöpfung, und sie allein ist wahrliche Kraft, Liebe, Wissen, Weisheit und Allmacht."}
    KossonSILVER0184 = Kanon {"Kanon  9.  16", "Und neben ihr ist keine andere Schöpfung in ihrem Himmel (Universum), und keine Götzen und keine Götter also."}
    KossonSILVER0185 = Kanon {"Kanon  9.  17", "Und denen, die die Schöpfung verleugnen; und denen, die andere Kräfte, Götter und Götzen neben sie setzen, kommet kein Zeichen von der Schöpfung zu ihnen."}
    KossonSILVER0186 = Kanon {"Kanon  9.  18", "Es sind die Schöpfungsabgewandten, die sehenden Blinden, die da nirgendwo sehen ein Zeichen ihrer Existenz, weder in allem Heer, das da lebet, noch im Inneren oder im Ausseren."}
    KossonSILVER0187 = Kanon {"Kanon  9.  19", "Es kommet zu ihnen auch nicht ein Zeichen der Wirklichkeit als Zeichen ihrer Wahrheit."}
    KossonSILVER0188 = Kanon {"Kanon  9.  20", "Und sie wandeln als Unwissende durch das Leben, fern der Wahrheit und dumm und blind."}
    KossonSILVER0189 = Kanon {"Kanon  9.  21", "Wissentlich wenden sie sich ab von der Wahrheit, denn ihr Leben des Lasters, der Gier und der Sucht, und ihr Leben der scheinbaren Mühelosigkeit, des Vergnügens, des Egoismus, der Selbstsucht und Machtgier ist ihnen näher als das Leben der Wahrheit und die Evolution zum SEIN(Schöpfungsleben)."}
    KossonSILVER0190 = Kanon {"Kanon  9.  22", "Sie wollen nicht sehen, dass sie sich abwenden vom wahrlichen Leben, das da führet zum SEIN （Schöpfungsleben) in alleiniger Form."}
    KossonSILVER0191 = Kanon {"Kanon  9.  23", "Sie verwerfen die Wahrheit, wie schon die Geschlechter vor ihnen, als ihnen die Wahrheit gebringet wurde."}
    KossonSILVER0192 = Kanon {"Kanon  9.  24", "Doch es ist ihnen Kunde gebracht von der Wahrheit, wie schon allen Geschlechtern vor ihnen."}
    KossonSILVER0193 = Kanon {"Kanon  9.  25", "Und wieder verwerfen und verleugnen sie die Wahrheit, und es wird ihnen keine Nachsicht und kein Verstehen zugesaget sein."}
    KossonSILVER0194 = Kanon {"Kanon  9.  26", "Und weiterhin beten (betteln und flehen) sie zu ihren Göttern und Götzen der Falschheit, un weiterhin bringen sie ihnen Verehrung entgegen."}
    KossonSILVER0195 = Kanon {"Kanon  9.  27", "Und weiterhin verwerfen und verleugnen sie die Schöpfung und damit die Wahrheit, und so die Liebe, das Wissen und die Weisheit also."}
    KossonSILVER0196 = Kanon {"Kanon  9.  28", "Ihnen ist aber neuerlich Kunde getan über die Wahrheit, die sie lachend verspotten."}
    KossonSILVER0197 = Kanon {"Kanon  9.  29", "Doch nicht unbestrafet bleibet ihr schändliches Tun."}
    KossonSILVER0198 = Kanon {"Kanon  9.  30", "Denn sie erlegen sich selbst Strafe auf und Vernichtung, so auch getan haben viele Geschlechter vor ihnen."}
    KossonSILVER0199 = Kanon {"Kanon  9.  31", "Sie selbst bestimmen sich die Strafe für ihr Unwesen und für ihre Verspottung und Verleugnung der Schöpfung und der Wahrheit und Liebe."}
    KossonSILVER0200 = Kanon {"Kanon  9.  32", "Ihre Strafe aber ruhet in frühen Toden und in Ängsten davor, und sie ruhet in Krankheit und Krieg, in Unfrieden und Hass, in Übellaunigkeit und Misserfolg, und in Schmerz, Zerstörung und Vernichtung."}
    KossonSILVER0201 = Kanon {"Kanon  9.  33", "Und sie erzeugen all die Dinge des Bösen in sich und um sich selbst."}
    KossonSILVER0202 = Kanon {"Kanon  9.  34", "Weil sie verleugnen die Schöpfung."}
    KossonSILVER0203 = Kanon {"Kanon  9.  35", "Und weil sie verleugnen die Wahrheit."}
    KossonSILVER0204 = Kanon {"Kanon  9.  36", "Und weil sie verleugnen die Liebe."}
    KossonSILVER0205 = Kanon {"Kanon 9.37-38", "Und weil sie verleugnen die Weisheit. Und weil sie verleugnen das Wissen."}
    KossonSILVER0206 = Kanon {"Kanon  9.  39", "Und weil sie verleugnen das Leben."}
    KossonSILVER0207 = Kanon {"Kanon  9.  40", "Und weil sie verleugnen das SEIN (Schöpfungsleben)."}
    KossonSILVER0208 = Kanon {"Kanon  9.  41", "Sie, die da fehlbar sind, sie sind Menschen, denen da ist gegeben auf der Erde ein Wohnsitz und alle Möglichkeit des Lebens und der Evolution zum SEIN (Schöpfungsleben)."}
    KossonSILVER0209 = Kanon {"Kanon  9.  42", "Doch sie gehorchen nicht den Gesetzen und Geboten der Schöpfung von alters her, obwohl ihnen gegeben ward die Wahrheit durch den JHWH und durch die Propheten."}
    KossonSILVER0210 = Kanon {"Kanon  9.  43", "Sie trieben Spott mit der Schöpfung, und mit ihren Gesetzen und Geboten und mit dem Leben also."}
    KossonSILVER0211 = Kanon {"Kanon  9.  44", "Sie hatten dadurch zu tragen die Folgen ihres verwerflichen Tuns."}
    KossonSILVER0212 = Kanon {"Kanon  9.  45", "Sie zerstöreten sich selbst, und sie vernichteten ihre Geschlechter."}
    KossonSILVER0213 = Kanon {"Kanon  9.  46", "Sie starben tausendfache und schmerzhafte Tode in vielfältigem Schreck und Entsetzen."}
    KossonSILVER0214 = Kanon {"Kanon  9.  47", "Und es wurden neue Geschlechter geboren, doch sie taten den Geschlechtern vor ihnen gleich, und auch diese vernichteten sich selbst."}
    KossonSILVER0215 = Kanon {"Kanon  9.  48", "Sie fanden Selbstvernichtung durch das falsche Leben, das sie führeten."}
    KossonSILVER0216 = Kanon {"Kanon  9.  49", "Und sie fanden Selbstvernichtung, indem sie verleugneten die Schöpfung und ihre Gesetze und Gebote."}
    KossonSILVER0217 = Kanon {"Kanon  9.  50", "Und wieder geschah Gleiches und Gleiches in endloser Folge, bis zum heutigen Tage."}
    KossonSILVER0218 = Kanon {"Kanon  9.  51", "Von alters her aber sandte ihnen hinab der JHWH der irdischen Menschengeschlechter die Schrift und das Wort der Wahrheit, auf Stein gehauen und geschrieben auf Tierhaut und auf Pergament, und es gab der Prophet das gleiche."}
    KossonSILVER0219 = Kanon {"Kanon  9.  52", "Und die Schrift und das Wort der Wahrheit waren niedergeschrieben durch der Propheten Beauftragte, und belehret durch ihn unter den Menschen der Erde, damit sie wissend würden."}
    KossonSILVER0220 = Kanon {"Kanon  9.  53", "Und sie konnten die Schrift des Wortes der Wahrheit um die Schöpfung und ihre Gesetze und Gebote fühlen mit ihren Händen."}
    KossonSILVER0221 = Kanon {"Kanon  9.  54", "Und sie konnten die Schrift lesen mit ihren Augen."}
    KossonSILVER0222 = Kanon {"Kanon  9.  55", "Und sie konnten die gesprochenen Worte der Schrift hören mit ihren Ohren."}
    KossonSILVER0223 = Kanon {"Kanon  9.  56", "Doch sie fühleten sie nicht, und sie sahen sie nicht, und sie höreten sie nicht."}
    KossonSILVER0224 = Kanon {"Kanon  9.  57", "Denn sie verschlossen ihre Sinne gegenüber der Wahrheit, und sie verschlossen all ihr Fühlen, und also verschlossen sie ihre Augen und ihre Ohren."}
    KossonSILVER0225 = Kanon {"Kanon  9.  58", "Und sie verleugneten die Schöpfung und ihre Gesetze und Gebote, und sie verleugneten den JHWH und den Propheten also."}
    KossonSILVER0226 = Kanon {"Kanon  9.  59", "Und sie blieben Unwissende und Verleugner der Wahrheit."}
    KossonSILVER0227 = Kanon {"Kanon  9.  60", "Und sie beschimpfeten den JHWH der Lüge und der Zauberei, und dem Propheten taten sie ebenso und trachteten ihm nach dem Leben."}
    KossonSILVER0228 = Kanon {"Kanon  9.  61", "Und sie beschimpfeten alle Propheten und verfolgeten sie, auf dass sie ihnen habhaft würden, und auf dass diese qualvoll des Todes sterben sollten."}
    KossonSILVER0229 = Kanon {"Kanon  9.  62", "Leben in Lästerlichkeit ward ihnen näher und erfüllender als wahrliches Leben zum SEIN （Schöpfungsleben)."}
    KossonSILVER0230 = Kanon {"Kanon  9.  63", "Und weiterhin lebeten sie in Lästerlichkeit in endloser Folge, bis auf den heutigen Tag."}
    KossonSILVER0231 = Kanon {"Kanon  9.  64", "Es muss daher in der neuen Zeit der JHWH durch den Propheten neuerlich die Schrift des wahrheitlichen Wortes hinabsenden."}
    KossonSILVER0232 = Kanon {"Kanon  9.  65", "Auf dass es vom Propheten der Neuzeit abermals gelehret, selbst niedergeschrieben und unter den Menschengeschlechtern belehret werde."}
    KossonSILVER0233 = Kanon {"Kanon  9.  66", "Auf dass in endlicher Folge der Mensch der Erde der Wahrheit trächtig werde und ihr folgsam sei."}
    KossonSILVER0234 = Kanon {"Kanon  9.  67", "Und auf dass der Prophet erkennet und die Lehre des Geistes Befolgung finden werde."}
    KossonSILVER0235 = Kanon {"Kanon  9.  68", "Dadurch werde die Sache endgültig und unwiderruflich letztmalig entschieden und in letzt licher Folge belehret, denn es gebühret dem Erdenmenschen keinerlei Aufschub mehr."}
    KossonSILVER0236 = Kanon {"Kanon  9.  69", "Im Namen der Schöpfung, es ist dies wahrheitlich und unverfälschlich und von edelster Form."}
    KossonSILVER0237 = Kanon {"Kanon  9.  70", "Gelobet sei die Wahrheit des Wortes der Wahrheit."}

    //Kanon 10
    KossonSILVER0238 = Kanon {"Kanon 10.   1", "Im Namen der Schöpfung, der Weisen, der Gerechten."}
    KossonSILVER0239 = Kanon {"Kanon 10.   2", "Preis (Ehre) sei der Schöpfung, der Allbarmherzigen."}
    KossonSILVER0240 = Kanon {"Kanon 10.   3", "Nicht hat der JHWH bestimmet einen seiner Helfer, der da wäre genennet beim Erdenmenschen Wächter oder Engel, um die Schrift der Belehrung zu bringen."}
    KossonSILVER0241 = Kanon {"Kanon 10.   4", "Es wäre Gefahr für sie im Leben, denn sie sind andersgeartet im Empfinden, und die Schwingung des Unwissens der Erdenmenschen würde sie töten."}
    KossonSILVER0242 = Kanon {"Kanon 10.   5", "Denn sie sind feinempfindsam, wie der JHWH also, und anfällig gegen die Schwingungen des Bösen und Negativen, das da dem Erdenmenschen noch ist eigen."}
    KossonSILVER0243 = Kanon {"Kanon 10.   6", "Und die Erdenmenschen würden sie auch anbeten (anbetteln und anflehen) und verehren, obwohl sie nicht sind die Schöpfung, der allein Anbetung (Ehrung) und Verehrung (Ehrwürdigung) gebühret."}
    KossonSILVER0244 = Kanon {"Kanon 10.   7", "Also sind erküret der JHWH und der Prophet, die da sind Mittler und Künder zwischen der Wahrheit und zwischen allen ihnen in eigener Pflichterkennung zugeordneten Menschengeschlechtern."}
    KossonSILVER0245 = Kanon {"Kanon 10.   8", "Und der Prophet ist Mensch der Erde, und seine Empfindsamkeit lieget an der Grenze des Wandelbaren, so er ist befähiget zu leben im Wellenbereiche der Erdenmenschen und doch sehen kann in die Güte des Anlitzes des JHWH."}
    KossonSILVER0246 = Kanon {"Kanon 10.   9", "Und der Prophet ist Mittler und Künder zwischen der Wahrheit und de Menschengeschlechtern, so er die Schrift des Wortes der Wahrheit zu bringen vermöge und die Belehrung also."}
    KossonSILVER0247 = Kanon {"Kanon 10.  10", "Und der Prophet kommet der Ebene des JHWH nahe, und er hat dessen gleiche Mission inne bei den Menschengeschlechtern."}
    KossonSILVER0248 = Kanon {"Kanon 10.  11", "Doch der Prophet ist Mensch, und der JHWH also."}
    KossonSILVER0249 = Kanon {"Kanon 10.  12", "Und über dem JHWH und über dem Propheten schwebet und lebet die Schöpfung als unmessbares Geheimnis."}
    KossonSILVER0250 = Kanon {"Kanon 10.  13", "Im Namen der Schöpfung, der Weisen, der Gerechten."}
    KossonSILVER0251 = Kanon {"Kanon 10.  14", "Preis (Ehre) sei der Schöpfung, die da bezeuget die Wahrheit des Wortes."}

    //Kanon 11
    KossonSILVER0252 = Kanon {"Kanon 11.   1", "Im Namen der Schöpfung, der Weisen, der Gerechten."}
    KossonSILVER0253 = Kanon {"Kanon 11.   2", "Preis (Ehre) sei der Schöpfung, die da ist wahrliche Liebe."}
    KossonSILVER0254 = Kanon {"Kanon 11.   3", "Wahrlich, der Prophet ist der Gesandte der Wahrheit."}
    KossonSILVER0255 = Kanon {"Kanon 11.   4", "Und über dem JHWH und dem Propheten schwebet die Kraft der Schöpfung als unmessbares Geheimnis."}
    KossonSILVER0256 = Kanon {"Kanon 11.   5", "Wahrlich, die Schöpfung und ihre Gesetze und Gebote werden vom Unwissenden verspottet."}
    KossonSILVER0257 = Kanon {"Kanon 11.   6", "Und wahrlich, es wird von den Unwissenden verspottet der JHWH."}
    KossonSILVER0258 = Kanon {"Kanon 11.   7", "Und wahrlich, es wird verspottet der Prophet von den Unwissenden."}
    KossonSILVER0259 = Kanon {"Kanon 11.   8", "Und wahrlich, was von den Unwissenden verspottet wird, das ist das allzeitig Existente rund um die Spötter."}
    KossonSILVER0260 = Kanon {"Kanon 11.   9", "Es werden die Spötter aber ihre gerechte Strafe finden, die sie sich selbst auferlegen und zur Wirkung bringen."}
    KossonSILVER0261 = Kanon {"Kanon 11.  10", "Der Erdenmensch wandle über seine Welt, und er wird sehen, wie das Ende derer ward, die da verspotteten die Wahrheit."}
    KossonSILVER0262 = Kanon {"Kanon 11.  11", "Und der Mensch wandle über seine Erde, und er wird sehen, wie das Ende derer ward, die da verleugneten die Schöpfung und ihre Gesetze und Gebote."}
    KossonSILVER0263 = Kanon {"Kanon 11.  12", "Und der Mensch wandle über seine Erde, und er wird sehen, wie das Ende derer ward, die da den JHWH verleugneten und den Propheten, und die den JHWH und die Propheten als Lügner behandelten."}
    KossonSILVER0264 = Kanon {"Kanon 11.  13", "Und wahrlich, der Mensch wandle über seine Erde, und er wird sehen und erkennen, wie das Ende derer ist in endloser Folge, die weiterhin den JHWH und den Propheten verleugnen, sie verspotten und der Lüge beschimpfen, und die da in Unwahrheit und Unwissenheit dem Leben und der Schöpfung und ihren Gesetzen und Geboten spotten."}
    KossonSILVER0265 = Kanon {"Kanon 11.  14", "Wehe allen, die da Spötter und Verleugner sind, denn ihrer harret schmerzvolles und angstvolles Sterben und auch solches Wiederleben."}
    KossonSILVER0266 = Kanon {"Kanon 11.  15", "Daran ist kein Zweifel, denn die Schöpfung lasset ihrer nicht spotten."}
    KossonSILVER0267 = Kanon {"Kanon 11.  16", "Ihre Gesetze und Gebote sind danach ausgerichtet, dass sich die Spötter, Verleugner und Unwissenden selbst in strengste Strafe legen, so ihnen Angst, Schmerz, Schreck und Entsetzen aus eigener Macht erblühen."}
    KossonSILVER0268 = Kanon {"Kanon 11.  17", "Die Schöpfung ist in allem Geheiligten(Ehrwürdigen)ihrer Kreationen, in allem,was da wohnet und existent ist, in der Nacht und am Tage, und sie ist die Allmächtige, die da kreieret hat alles in Gesetze und Gebote, und in alles, was da kreuchet und fleuchet."}
    KossonSILVER0269 = Kanon {"Kanon 11.  18", "Und wahrlich, der Mensch soll niemals und keine andere Schutzkraft neben die Schöpfung stellen als die Schöpfung selbst, die da gebar Himmel und Erde, und die da allem Kreuchenden und Fleuchenden Nahrung gebet und selbst keiner ähnlichen Nahrung bedarf."}
    KossonSILVER0270 = Kanon {"Kanon 11.  19", "Und es ist geboten, dass der Mensch und alles Leben den schöpferischen Gesetzen und Geboten folge, so alles Leben ein schöpfungsgefälliges Dasein führe."}
    KossonSILVER0271 = Kanon {"Kanon 11.  20", "Und es sei kein Mensch und keinerlei andere Lebensform, die der Schöpfung Götter und Götzen oder andere Nebenbuhler an die Seite stelle."}
    KossonSILVER0272 = Kanon {"Kanon 11.  21", "Und sollte dies doch getan werden in Ungehorsam, dann fürchte der Fehlbare den Tag der furchtbaren Strafe, die da entstehend ist aus seinem eigenen Antrieb und aus dem Leben des Fehlbaren selbst heraus, sowohl im Ausmasse als auch in der Furchtbarkeit der Strafe."}
    KossonSILVER0273 = Kanon {"Kanon 11.  22", "Wer sich aber vor Strafe bewahret und lebet getreu den Gesetzen und Geboten der Schöpfung, dem offenbaret sich Glückseligkeit."}
    KossonSILVER0274 = Kanon {"Kanon 11.  23", "Wer sich aber Strafe auferleget durch Spott und Verleugnung und Unwissenheit, der verfallet in Unglückseligkeit, und keiner wird sein rund um ihn, der sie von ihm hinwegzunehmen vermag."}
    KossonSILVER0275 = Kanon {"Kanon 11.  24", "Nur die Befolgung der Wahrheit vermag die Strafe und die Unglückseligkeit zu mildern und hinwegzunehmen."}
    KossonSILVER0276 = Kanon {"Kanon 11.  25", "Und der Weg der Erlösung führet allein über die Befolgung der Gesetze und Gebote der Schöpfung."}
    KossonSILVER0277 = Kanon {"Kanon 11.  26", "Wird der Mensch berühret von Glück oder Unglück, dann ist dies gelegen allein in seinem eigenen Tun, je nachdem, ob er die schöpferischen Gesetze und Gebote befolget oder nicht."}
    KossonSILVER0278 = Kanon {"Kanon 11.  27", "Es hat der Mensch die Kraft, alles zu tun, was er will, so im Guten und im Bösen, so ihn Glück oder Unglück berühret je nach seinem Tun."}
    KossonSILVER0279 = Kanon {"Kanon 11.  28", "Der Mensch selbst ist die Macht über sich selbst."}
    KossonSILVER0280 = Kanon {"Kanon 11.  29", "Und nutzet er die Macht seiner selbst zum Bösen, dann ist er ein Narr."}
    KossonSILVER0281 = Kanon {"Kanon 11.  30", "Und nutzet er die Macht seiner selbst zum Guten, dann ist er ein Weiser."}
    KossonSILVER0282 = Kanon {"Kanon 11.  31", "Und befolget er die Gesetze und Gebote der Schöpfung, dann ist er der Kundige und der Könnende."}
    KossonSILVER0283 = Kanon {"Kanon 11.  32", "Und es ist dies die Wahrheit der Wahrheit, die da zeuget von sich selbst."}
    KossonSILVER0284 = Kanon {"Kanon 11.  33", "Und welches Ding ist wahrlich am gewichtigsten als Zeugnis?"}
    KossonSILVER0285 = Kanon {"Kanon 11.  34", "Wahrlich, allein die Wahrheit ist das wahrliche Zeugnis der Wahrheit."}
    KossonSILVER0286 = Kanon {"Kanon 11.  35", "Und alles dies ist das Zeugnis der Wahrheit, so also die Schrift des Wortes der Wahrheit, auf dass alle Menschengeschlechter durch dieses Zeugnis belehret und gewarnet seien, wo und wann sie das Zeugnis der Wahrheit auch immer erreichen wird."}

    //Kanon 12
    KossonSILVER0287 = Kanon {"Kanon 12.   1", "Im Namen der Schöpfung, der Weisen, der Gerechten."}
    KossonSILVER0288 = Kanon {"Kanon 12.   2", "Preis (Ehre) sei der Schöpfung, der Wahrheitlichen."}
    KossonSILVER0289 = Kanon {"Kanon 12.   3", "Wahrlich, es ist gegeben durch Irrlehren, Selbstsucht und Machtgier der Menschen, dass sie beweisen und bezeugen wollen, dass neben der Schöpfung andere Kräfte und Götzen und Götter seien."}
    KossonSILVER0290 = Kanon {"Kanon 12.   4", "Nicht ist es in Wahrheit aber so, und jegliches diesartige Zeugnis ist falsch, verlogen und unbeweisbar."}
    KossonSILVER0291 = Kanon {"Kanon 12.   5", "Zu bezeugen ist aber, dass die Schöpfung die alleinige und einzige Schöpfung ist in ihrem Himmel (Universum), allzeitlich und ewiglich."}
    KossonSILVER0292 = Kanon {"Kanon 12.   6", "Und wahrlich, es ist ihr alles fern, unwürdig und erleugnet, was ihr Menschen ihr fälschlich andichtet und zur Seite stellet."}
    KossonSILVER0293 = Kanon {"Kanon 12.   7", "Der JHWH aber weiss wahrlich um die Wahrheit, und wahrlich ist ebenso wissend um diese Wahrheit der Prophet."}
    KossonSILVER0294 = Kanon {"Kanon 12.   8", "Und wissend sollen alle jene werden, die durch das Buch des Wortes der Wahrheit und durch den Propheten belehret werden, so aus ihnen Geistführer entstehen, die den Menschengeschlechtern belehrend sein sollen."}
    KossonSILVER0295 = Kanon {"Kanon 12.   9", "Jene aber, die ihr Leben verderben, sie sind Unwissende und Zweifler, Spötter, Verleugner, Übellaunige, Unzufriedene, Ungerechte und die stetigen wahrlichen Verlierer, und sie sollen da unterrichtet und belehret werden in der Wahrheit."}
    KossonSILVER0296 = Kanon {"Kanon 12.  10", "Sie sollen unterrichtet und belehret werden, auf dass sie Wissende werden und auf dass sie die Schöpfung ehren in Ehrfurcht (Ehrung, Ehrerbietung, Ehrwürdigung, Ehrwürdigkeit) und in Erkennung der Wahrheit."}
    KossonSILVER0297 = Kanon {"Kanon 12.  11", "Und sie sollen belehret werden, auf dass sie den Weg des Verderbens verlassen und auf dass sie die Gesetze und Gebote der Schöpfung erfüllen, so sie ein schöpfungsgerechtes Leben führen, und ihre nachfolgenden Geschlechter also."}
    KossonSILVER0298 = Kanon {"Kanon 12.  12", "Und wer bedarf der Wahrheit mehr als die Unwissenden und die Fehlbaren, und die Spötter, Verderber und Verleugner, die da sind die wahrlichen und stetigen Verlierer!"}
    KossonSILVER0299 = Kanon {"Kanon 12.  13", "Denn wer ist ungerechter als sie, die sie Lügen ersinnen wider die Schöpfung, wider ihre Gesetze und Gebote, und wider den JHWH und wider den Propheten, und wider das Leben und wider das SEIN (Schöpfungsleben)."}
    KossonSILVER0300 = Kanon {"Kanon 12.  14", "Wo aber ist die Weisheit, die Liebe, die Wahrheit und das Wissen dieser Ungerechten und Lügner, und wo ist ihre Glückseligkeit?"}
    KossonSILVER0301 = Kanon {"Kanon 12.  15", "Sie sind die Verführten, die sich selbst verführet haben und die allzeitlich die stetigen Verlierer sind, wenn sie den Pfad des Verderbens nicht verlassen."}
    KossonSILVER0302 = Kanon {"Kanon 12.  16", "Und sie sind die Verlierer am Tage ihres selbsterwählten Gerichts, wenn sie da scheinheilig （scheinehrwürdig/falsch), verlogen und selbstmitleidig mit Narrenworten Rechenschaft suchen und sprechen, dass sie in gutem Glauben ihre Götter und Götzen erschaffen und erhoben und angebetet (angebettelt und angefleht) hätten."}
    KossonSILVER0303 = Kanon {"Kanon 12.  17", "Und sie werden rechenschaftsfordernd sprechen, dass ihnen nur Lüge und Trug, jedoch niemals die Wahrheit belehret worden sei, doch es wird niemand sein, der ihren Worten lauschet und ihren irren Glauben als Wahrheit nennet."}
    KossonSILVER0304 = Kanon {"Kanon 12.  18", "Denn mitnichten können sie Rechenschaft ablegen und ihren Glauben der Wahrheit gleichstellen und ihre verlogene Rede führen, denn ihnen und allen Geschlechtern zuvor ward die Wahrheit kundgetan."}
    KossonSILVER0305 = Kanon {"Kanon 12.  19", "Und so ist die Wahrheit auch kundgetan in der Neuzeit, durch den Propheten dieser Zeit, der da ist wirkend dem JHWH gleich auf Erden und der gebend ist die Schrift des Wortes der Wahrheit andermalig."}
    KossonSILVER0306 = Kanon {"Kanon 12.  20", "Und gegeben ist die Schrift des Wortes der Wahrheit ein andermal und ein letztes Mal, damit die Lehre gegeben sei in endgültiger und unverfälschter Form für alle kommende Zeit."}
    KossonSILVER0307 = Kanon {"Kanon 12.  21", "Auf dass der Erdenmensch endlich erkenne die Wahrheit und sie befolge fortan."}

    //Kanon 13
    KossonSILVER0308 = Kanon {"Kanon 13.   1", "Im Namen der Schöpfung, der Weisen, der Gerechten."}
    KossonSILVER0309 = Kanon {"Kanon 13.   2", "Preis (Ehre) sei der Schöpfung, die da ist Gerechtigkeit."}
    KossonSILVER0310 = Kanon {"Kanon 13.   3", "Und so spricht der JHWH: «Wahrlich, die, welche das Buch des Wortes der Wahrheit als solches erkennen, die erkennen es so, wie sie ihre eigenen Kinder erkennen."}
    KossonSILVER0311 = Kanon {"Kanon 13.   4", "Jene aber, welche die Wahrheit des Wortes nicht erkennen, die verderben ihr Leben, und sie sind Ungerechte und Lügner an der Sache des Lebens."}
    KossonSILVER0312 = Kanon {"Kanon 13.   5", "Wer aber ist ungerechter als die Ungerechten, die da Lüge ersinnen wider die Schöpfung, wider den JHWH und wider den Propheten, und deren Lehre sie der Lüge zeihen."}
    KossonSILVER0313 = Kanon {"Kanon 13.   6", "Wahrlich, diese Ungerechten schlagen sich bewusst selbst mit Blindheit, und sie erlangen keine Glückseligkeit, und sie strafen sich selbst Lügner also."}
    KossonSILVER0314 = Kanon {"Kanon 13.   7", "Und am Tage ihres selbstbestimmten Gerichts wird an sie die Frage ergehen, wo nun die Hilfe und Existenz ihrer selbsterdachten Götter und Götzen sei, wenn sie vor dem Anlitz des Todes stehen und in Angst und Furcht erzittern."}
    KossonSILVER0315 = Kanon {"Kanon 13.   8", "Und es wird dann sein bei ihnen Heulen und Zähneklappern."}
    KossonSILVER0316 = Kanon {"Kanon 13.   9", "Und ihre Rede wird kein Gehör finden, wenn sie sich da rechtfertigen und sagen, dass sie keine Götterdiener und Götzendiener gewesen seien."}
    KossonSILVER0317 = Kanon {"Kanon 13.  10", "Ihre Rechenschaft wird nur Lüge sein, und sie werden wider sich selbst lügen, und sie werden lügen wider all das, was sie sich an Göttern und Götzen ersannen, die ihnen am Tage ihres selbsternannten Gerichtes weder Angst noch Furcht nehmen, noch ihnen hilfreich sind."}
    KossonSILVER0318 = Kanon {"Kanon 13.  11", "Der Tag und die Stunde des selbsternannten Gerichts wird für die Ungerechten und Lügner, Spötter und Verleumder Schrecknis und Entsetzen sein, und sie werden tausendfache Tode büssen und sich selbst strafen für die Buhlerei um ihre selbsternannten Götter und Götzen."}
    KossonSILVER0319 = Kanon {"Kanon 13.  12", "Und die Schuld lieget bei den Ungerechten allein, und ihre Strafe und ihr Entsetzen also."}
    KossonSILVER0320 = Kanon {"Kanon 13.  13", "Und viele sind unter den Ungerechten, die da dem Wort der Wahrheit und der Lehre Gehör schenken, gegen das Gehörte jedoch taub bleiben, so sie nicht verstehen können, weil sie die Wahrheit nicht verstehen wollen."}
    KossonSILVER0321 = Kanon {"Kanon 13.  14", "Und viele sind unter den Ungerechten, die da ein jedes Zeichen der Wahrheit und der Schöpfung und ihrer Gesetze und Gebote sehen, doch sie nehmen sie nicht wissend und in Erkenntnis in sich auf, weil sie die wahrliche Wahrheit nicht sehen wollen."}
    KossonSILVER0322 = Kanon {"Kanon 13.  15", "Die Ungerechten verleugnen die Wahrheit, und sie ersinnen Lügen wider die Wahrheit, um Streit zu finden mit den Gerechten und Wissenden und um lügen zu können, dass alles nur Fabelmachenschaft sei von alters her, und Zauberei und Betrug des JHWH und des Propheten also."}
    KossonSILVER0323 = Kanon {"Kanon 13.  16", "Und die Ungerechten und Nichtrechtschaffenen halten sich von der Lehre der Wahrheit fern, oder sie erlernen sie nur, um sie zu verunglimpfen."}
    KossonSILVER0324 = Kanon {"Kanon 13.  17", "Und sie erfinden Lügen wider die Wahrheitsbringer, wider den JHWH und wider den Propheten, um andere von der Lehre des Geistes und der Wahrheit und von der allumfassenden Liebe und Güte der Schöpfung abzuhalten."}
    KossonSILVER0325 = Kanon {"Kanon 13.  18", "Sie beschimpfen den Propheten der Lüge, des Schwindels, des Betrugs, der Zauberei und der Magie, so aber auch der Taschenspielerei und der Profitgier also."}
    KossonSILVER0326 = Kanon {"Kanon 13.  19", "Und sie beschimpfen die durch den Propheten dargebrachte Lehre des Geistes als Mummenschanz und abartige Phantasie, und die Ungerechten bemühen sich, aus ihren derartigen Lügen Profit zu schlagen und Macht und Befriedigung also."}
    KossonSILVER0327 = Kanon {"Kanon 13.  20", "Sie verbieten den nach Wahrheit Suchenden und Hoffenden die Wahrheit zu finden und nach den Gesetzen und Geboten der Schöpfung zu leben, und sie verbieten ihnen, den wahrheitlichen Worten des Propheten und der Wahrheit Ohr zu sein."}
    KossonSILVER0328 = Kanon {"Kanon 13.  21", "Und die Ungerechten halten sich selbst von der Wahrheit fern, und vor dem Propheten und dessen wahrheitlichem Wort und dessen Ratgebung also."}
    KossonSILVER0329 = Kanon {"Kanon 13.  22", "Wahrlich, sie stürzen sich aber mit ihrem Tun und Denken ins eigene Verderben, und die von ihnen mitgerissen werden und die ihnen hörig sind also."}
    KossonSILVER0330 = Kanon {"Kanon 13.  23", "Allein, sie alle begreifen es nicht, weil sie es nicht begreifen wollen, denn der eigene Profit und das Leben der mühelosen Lästerlichkeit ist ihnen näher als das Leben der Glückseligkeit, das erarbeitet werden muss mit Mühen."}
    KossonSILVER0331 = Kanon {"Kanon 13.  24", "Sie alle durchschreiten das grosse Tor zur breiten Strasse, die da das Leben der Lästerlichkeit bietet."}
    KossonSILVER0332 = Kanon {"Kanon 13.  25", "Sie vermeiden die kleine und enge Pforte, der ein schmaler und steiniger Weg folget und der einige Bemühung der Begehung erfordert, die aber von ihnen vermieden wird, denn Mühe, so denken sie, möge sich nicht lohnen."}
    KossonSILVER0333 = Kanon {"Kanon 13.  26", "Wenn sie dann jedoch stehen am Ende ihrer Tage, wenn der Tod vor ihr Anlitz tritt, dann werden sie reuig und wünschen, ihr Leben ein andermal zu leben."}
    KossonSILVER0334 = Kanon {"Kanon 13.  27", "Nun heucheln sie plötzlich erkennend zu sein in jenen Dingen, die sie ihr Leben lang haben verleugnet, verspottet und ungerecht behandelt, und sie wünschen in ihr gelebtes Leben zurückgebracht zu werden, um es ein andermal zu leben, und in nun gerechter Weise."}
    KossonSILVER0335 = Kanon {"Kanon 13.  28", "Ihre Wünsche in Angst, Furcht, Schrecken und Entsetzen jedoch verhallen ungehöret, denn niemand vermag ihnen Ohr zu sein, und Hilfe bleibet aus also, so sie in Angst und Schrecken vom Diesseits ins Jenseits wechseln, belastet mit Angst und Furcht für das nächste Wiederleben."}
    KossonSILVER0336 = Kanon {"Kanon 13.  29", "Ach, hätten sie die Worte und die Zeichen der Wahrheit nicht als Lüge behandelt, und hätten sie nicht den Propheten verfolget, es wäre ihnen das Sterben keine Last und also auch nicht das Leben."}
    KossonSILVER0337 = Kanon {"Kanon 13.  30", "Im Angesicht des Todes zumindest wird den Ungerechten in Angst und Schrecken klar, was sie bewusst widerrechtlich ehemals zu verhehlen pflegeten."}
    KossonSILVER0338 = Kanon {"Kanon 13.  31", "Und wahrlich, wären gestaltet die Gesetze und Gebote der Schöpfung dergleichen, dass sie vor dem Anlitz des Todes zurückweichen und abermalig ihr vertanes Leben neuerlich leben könnten, dann wäre ganz gewiss, dass sie bald zu dem ihnen Verbotenen zurückkehren würden."}
    KossonSILVER0339 = Kanon {"Kanon 13.  32", "Abermals würden sie ein Leben in Lästerlichkeit führen, in Spötterei und Verleugnung, und neuerlich würden sie naschen und essen von den Früchten des Verbotenen."}
    KossonSILVER0340 = Kanon {"Kanon 13.  33", "Also sind sie noch Lügner auf dem Lager des Sterbens und im Angesichte des Todes, wenn sie da heulend und zähneklappernd um Gnade winseln, die von niemandem kann gegeben sein."}
    KossonSILVER0341 = Kanon {"Kanon 13.  34", "Und ihr gnadeerflehendes Winseln ist nur Lug und Trug, wie ihr Leben ehedem, denn ihr Schreien in Angst und Furcht ist nur Schein und Betrug und momentan, um noch letztlich am unausweichlichen Ende Profit für sich zu heischen."}
    KossonSILVER0342 = Kanon {"Kanon 13.  35", "Also handeln sie selbst zur letzten Spanne ihrer Lebenszeit eigensüchtig und profitbedacht zum eigenen Vorteil, ohne Ehrlichkeit und Ehrfurcht (Ehrung, Ehrerbietung, Ehrwürdigkeit, Ehrwürdigung) zur Schöpfung und zur Wahrheit."}
    KossonSILVER0343 = Kanon {"Kanon 13.  36", "Wie also könnte da erfolgen Gnade oder Nachsicht, Aufschub und neuerliche Gelegenheit zu leben ein zweites Mal dasselbe Leben."}
    KossonSILVER0344 = Kanon {"Kanon 13.  37", "Wahrlich, sie werden sein die lebenslangen Verlierer, denn nicht unbestrafet leugnen und verleugnen sie die Schöpfung und die Wahrheit."}
    KossonSILVER0345 = Kanon {"Kanon 13.  38", "Und nicht ungestrafet verleugnen und beschimpfen sie den JHWH und den Propheten."}
    KossonSILVER0346 = Kanon {"Kanon 13.  39", "Und nicht ungestrafet verleugnen sie des Propheten Begegnung mit dem JHWH."}
    KossonSILVER0347 = Kanon {"Kanon 13.  40", "Denn sie werden sich selbst bestrafen für ihre Lügen und Verleumdungen, wenn für sie unverhoffet die Stunde des von ihnen selbstbestimmten Gerichtes über sie kommet, und so beiihnen wird sein Heulen und Zähneklappern."}
    KossonSILVER0348 = Kanon {"Kanon 13.  41", "Und wehe ihnen, die sie das Leben vernachlässigen, denn sie werden Mühsal, Schmerz und Not erdulden müssen, und ihre Lasten werden sie auf ihren Rücken tragen müssen, schwer und niederdrückend."}
    KossonSILVER0349 = Kanon {"Kanon 13.  42", "Wahrlich, schlimm ist das, womit sich die Ungerechten beladen, denn ihre Last ist gewichtiger, als dass sie diese erkennen und tragen könnten."}
    KossonSILVER0350 = Kanon {"Kanon 13.  43", "Wehe ihnen, die sie da sind Ungerechte, Lügner, Verleumder, Spötter und Vernachlässiger."}
    KossonSILVER0351 = Kanon {"Kanon 13.  44", "Ihrer harret selbsterschaffenes Übel, Angst und Furcht, und Schrecken und Entsetzen also."}
    KossonSILVER0352 = Kanon {"Kanon 13.  45", "Und ihrer harret böses selbsterzeugetes Gericht, so im Leben als auch dann, wenn sie betreten die Scheide zwischen Diesseits und Jenseits, und wenn sie stehen im Wiederleben also."}
    KossonSILVER0353 = Kanon {"Kanon 13.  46", "Glücklich sind aber die Gerechten, die da haben gelebet in Wahrheit, denn ihrer ist die Glückseligkeit im Leben, und ihrer ist die Freude des Todes also,  wenn sie in dessen Anlitz blicken."}
    KossonSILVER0354 = Kanon {"Kanon 13.  47", "Nicht wird ihnen sein Angst und Furcht, und es wird für sie sein weder Schrecken noch Entsetzen, noch unzeitiges Ende, denn der Gerechte ist wissend und weise, so ihm das Ende neuer Anfang ist zum Wiederleben."}
    KossonSILVER0355 = Kanon {"Kanon 13.  48", "Im Namen der Schöpfung, die da ist Gerechtigkeit, dies entsprechet dem wahrheitlichen Gang."}
    KossonSILVER0356 = Kanon {"Kanon 13.  49", "Preis (Ehre) sei der Schöpfung, die da alles hat erschaffen in Liebe, in Weisheit, in Wahrheit, und wissend in Gerechtigkeit.»"}

    //Kanon 14
    KossonSILVER0357 = Kanon {"Kanon 14.   1", "Im Namen der Schöpfung, der Weisen, der Gerechten."}
    KossonSILVER0358 = Kanon {"Kanon 14.   2", "Preis (Ehre) sei der Schöpfung, die da ist Wahrlichkeit."}
    KossonSILVER0359 = Kanon {"Kanon 14.   3", "Und es sprichet der JHWH zum Propheten als Belehrung für die irdischen Menschengeschlechter:"}
    KossonSILVER0360 = Kanon {"Kanon 14.   4", "«Nicht ist das Leben nur ein Spiel und ein Zeitvertreib, denn wahrlich dienet das Leben den Rechtschaffenen als Zeit des Lernens und der Evolution."}
    KossonSILVER0361 = Kanon {"Kanon 14.   5", "Und wahrlich, dies ist diesseitig so auf Erden als auch anderseitig im Bereiche des Jenseits der Erde."}
    KossonSILVER0362 = Kanon {"Kanon 14.   6", "Diesseits der Erde ist das Leben des grobstofflich Materiellen und die damit verbundene Lernzeit und Evolution."}
    KossonSILVER0363 = Kanon {"Kanon 14.   7", "Im jenseitigen Raume aber ist das Leben des Geistigen und die damit verbundene Evolution der Geistform, und also findet statt im jenseitigen Raume die Vorbereitung der Geistform auf das Wiederleben in einem neuen Menschenkörper."}
    KossonSILVER0364 = Kanon {"Kanon 14.   8", "Und das Diesseits und das Jenseits sind im selbigen Raum gleichsam, getrennet jedoch in zwei ungleichsame Dimensionen, so in die sichtbare und greifbare materielle Ebene, und in die unsichtbare und ungreifbare geistige Ebene."}
    KossonSILVER0365 = Kanon {"Kanon 14.   9", "Also sind gegeben zwei ungleichsame Dimensionen, die materielle Ebene, in der das physische Leben sein Dasein pfleget, und die geistige Ebene, in die die Geistform des Menschen einschwebet nach dem Sterben des materiellen Körpers, um im Jenseitsraume zu leben."}
    KossonSILVER0366 = Kanon {"Kanon 14.  10", "Also ist der Jenseitsraum der Ort also, in den die Geistform des Menschen eintretet, so er den materiellen Körper des Menschen ableget durch den Tod."}
    KossonSILVER0367 = Kanon {"Kanon 14.  11", "Und der Tod des Menschen ist einem Tiefschlafe gleich, erholsam und entwickelnd im unsichtbaren Jenseitsraume für die Geistform, durch die im physischen Leben des Menschen der Körper Belebung findet."}
    KossonSILVER0368 = Kanon {"Kanon 14.  12", "Und die Geistform des Menschen, die da ist von Winzigkeit wie ein Floh, entwickelt sich im Körper des Menschen in Liebe, Wissen, Weisheit und Wahrheit, so aber sie sich auch evolutionieret in erweiternder Form im Jenseitsbereiche also."}
    KossonSILVER0369 = Kanon {"Kanon 14.  13", "Und für die Geistform folget im Jenseitsbereiche Evolution und Vorbereitung für das nächste Wiederleben, so sie neuerlich geboren werde in einem materiellen Menschenkörper unbeschweret und frisch als neue Persönlichkeit, und auf der Erde ein andermal."}
    KossonSILVER0370 = Kanon {"Kanon 14.  14", "Und es lebet der Mensch im Diesseits auf Erden, und er gehet in den Tod ins Jenseits, so er unaufhaltsam wiedergeboren werde und evolutionieret."}
    KossonSILVER0371 = Kanon {"Kanon 14.  15", "Also lebet der Mensch und sterbet und wird wiedergeboren als neue Persönlichkeit in neuem Körper, und es wiederholet sich das diesartige Werden und Vergehen so lange, bis die Zeit der Zeit angebrochen ist, zu der die Geistform des Menschen als winziges Teilstück der Schöpfung zur weiteren Evolution keines materiellen Menschenkörpers mehr bedarf."}
    KossonSILVER0372 = Kanon {"Kanon 14.  16", "Also lebet und evolutionieret die Geistform fortan ohne materiellen Menschenkörper, denn wahrlich lebet und evolutionieret sie dann in reiner geistiger Form und als feinstofflicher Geistkörper."}
    KossonSILVER0373 = Kanon {"Kanon 14.  17", "Und es dauert dieses Werden und Vergehen des Menschen und seiner Geistform im Masse der Zeit 70 Milliarden Erdenjahre, errechnet vom Anbeginn der Existenz des Menschen in jeweiliger Person, bis hin zur Lichtwerdung der Geistform ohne materiellen Körper."}
    KossonSILVER0374 = Kanon {"Kanon 14.  18", "Also ist gegeben durch die Gesetze und Gebote der Schöpfung, dass sich in wiederholender Folge des Werdens und Vergehens, des Lebens und des Todes und der Wiedergeburt die Evolution des Menschen erfülle, und die Evolution der Geistform des Menschen also."}
    KossonSILVER0375 = Kanon {"Kanon 14.  19", "Es lebet der Mensch also in vielfacher Folge, eingeordnet in Geburt und Leben, und in Tod und Wiedergeburt, so gegeben ist die unaufhaltsame Evolution und die Erfüllung des Zieles zum SEIN (Schöpfungsleben)."}
    KossonSILVER0376 = Kanon {"Kanon 14.  20", "Und gesaget ist dadurch, dass der Mensch nicht lebet nur einmal, sondern dass er lebet wieder und wieder, durch Geburt und Sterben und Wiedergeburt, die da sind in endloser Folge."}
    KossonSILVER0377 = Kanon {"Kanon 14.  21", "Und wahrlich, es ist dadurch gegeben, dass der Mensch im Leben nach seinem Leben und so also im nächsten Leben die Früchte dessen erntet, was er hat gepflanzet im Leben zuvor, weshalb er stetig sei darauf bedacht, sein Leben zu führen in schöpfungsgefälliger Form, und als rechtschaffener Gerechter gleichsam.»"}

    //Kanon 15
    KossonSILVER0378 = Kanon {"Kanon 15.   1", "Im Namen der Schöpfung, der Weisen, der Gerechten."}
    KossonSILVER0379 = Kanon {"Kanon 15.   2", "Preis (Ehre) sei der Schöpfung, der Geduldigen."}
    KossonSILVER0380 = Kanon {"Kanon 15.   3", "Und also spricht der JHWH zum Propheten, der da ist betrübet in seiner Mission:"}
    KossonSILVER0381 = Kanon {"Kanon 15.   4", "«Siehe, du bist betrübet in deinen Gedanken und in deinem Fühlen, denn die Menschen zeihen dich der Lüge und der Falschheit, und den JHWH zeihen sie so also."}
    KossonSILVER0382 = Kanon {"Kanon 15.   5", "Und sie zeihen die Schöpfung der Falschheit, und ihre Gesetze und Gebote also."}
    KossonSILVER0383 = Kanon {"Kanon 15.   6", "Jedoch die Zeihenden sind Frevler und Ungerechte unter den Menschen, denn sie verwerfen nicht nur die Begegnung des JHWH mit dem Propheten, sondern die Existenz der Schöpfung ebenso, denn sie setzen der Schöpfung einen Gewaltherrscher gleich (bei), den sie benennen als Gott."}
    KossonSILVER0384 = Kanon {"Kanon 15.   7", "Und sie verwerfen die Zeichen der Schöpfung, die da sind gegeben Tag um Tag und Nacht um Nacht, in jeglichem Lebendigen und Existenten, so in allem Heer des Lebens, das da ist kreieret durch die Schöpfung."}
    KossonSILVER0385 = Kanon {"Kanon 15.   8", "Und sie beschimpfen dich, den Propheten und Gesandten der Wahrheit, und sie beschimpfen dich der Falschheit schon von alters her, der du schon warst der Prophet in vielfacher Folge vor der Neuzeit."}
    KossonSILVER0386 = Kanon {"Kanon 15.   9", "Aber nicht betrübet es dich, dass sie dich als Prophet beschimpfen und dass sie dich der Falschheit verleugnen, denn Betrübtheit brennet in dir, weil verleugnet wird die Lehre des Geistes, und so das Wort der Wahrheit und die Schöpfung mit ihren Gesetzen und Geboten."}
    KossonSILVER0387 = Kanon {"Kanon 15.  10", "Wahrlich, du bist geduldig, und Geduld ist dir im Grossen eigen."}
    KossonSILVER0388 = Kanon {"Kanon 15.  11", "Deine Geduld aber erhellet nicht deine Betrübtheit, und sie lichtet nicht die dunklen Schleier des in dir in Liebe erwalleten heiligen (kontrollierten) Zornes."}
    KossonSILVER0389 = Kanon {"Kanon 15.  12", "Du bist von den Menschen verleugnet, so wie ehedem dir angetan ward zu früheren Zeiten deiner früheren Leben, als du warst in Mission in nämlicher Form für die Wahrheit und da du warst verfolget und geschunden oft und oft durch die Ungerechten."}
    KossonSILVER0390 = Kanon {"Kanon 15.  13", "Gegeben waren aber bei den irdischen Menschengeschlechtern noch andere Propheten, die da waren Gesandte der Wahrheit ebenso."}
    KossonSILVER0391 = Kanon {"Kanon 15.  14", "Und sie waren betrübet also, und gescholten und geschunden und beschimpfet als lügenhaft und falsch, wie dir geschah und neuerlich geschehet also."}
    KossonSILVER0392 = Kanon {"Kanon 15.  15", "Und so dir geschah und geschehet und drohet, so waren auch die anderen Propheten verfolget und geschunden ohne Unterlass, und ihr Lohn für Liebe und Wahrheit war gewaltsamer Tod durch Meuchlerhand, so dir war eigen also."}
    KossonSILVER0393 = Kanon {"Kanon 15.  16", "Doch es waren welche unter ihnen, unter den Propheten, die da in Geduld waren kleinmütig und schwach, oder die da waren falsche Propheten in eigener Gnade, dem Gleichsames sich zutraget bis auf den heutigen Tag, so sich der Mensch bewahre vor falschen Propheten."}
    KossonSILVER0394 = Kanon {"Kanon 15.  17", "Die da aber waren wahrliche Propheten und gesendet durch eigenes Pflichtgefühl für die Wahrheit, gleichsam dir; sie blieben geduldig und wahrheitlich, und sie trotzeten aller Beschuldigung und aller Verfolgung."}
    KossonSILVER0395 = Kanon {"Kanon 15.  18", "Und wahrlich, auch du trotzest den Verleugnern und Verfolgern und Beschuldigern, und so also den Spöttern und Lügnern, so du getan hast ehedem."}
    KossonSILVER0396 = Kanon {"Kanon 15.  19", "Und du trotzest aller Unbill und aller Schande, die ungerecht über dich gebracht wird, und du trotzest auch allen Feindanschlägen, die dir harmend sein sollen und vernichtend für die Mission."}
    KossonSILVER0397 = Kanon {"Kanon 15.  20", "Und siehe, du bist betrübet über diese Dinge der Ungerechten, doch hast du zu bedenken, dass deine Betrübtheit in dir die Geduld erweitert, die da, mit wahrlicher Liebe verbündet, dich hoffen und harren lasset."}
    KossonSILVER0398 = Kanon {"Kanon 15.  21", "Und dein Hoffen und Harren ist gerichtet darauf, dass die Wahrheit den Sieg erringe und dass die irdischen Menschengeschlechter den Pfad der Wahrheit zur Schöpfung beschreiten werden."}
    KossonSILVER0399 = Kanon {"Kanon 15.  22", "Also weisst du: Weder auf der Erde noch anderswo gibt es einen Menschen, der das Wort der Wahrheit und die Gesetze und Gebote der Schöpfung zu ändern vermöchte."}
    KossonSILVER0400 = Kanon {"Kanon 15.  23", "Du bist dieser Wahrheit wissend seit alters her, so du danach lebest und belehrest."}
    KossonSILVER0401 = Kanon {"Kanon 15.  24", "Also ist dir eigen die Wahrheit und die Weisheit, und so auch das Wissen und die Liebe, weshalb du dein Haupt nicht beugen sollst in Gram."}
    KossonSILVER0402 = Kanon {"Kanon 15.  25", "Du bist der Prophet der Neuzeit, und beständig in dieser Mission seit alters her, und du bist wahrlich der Gesandte der Wahrheit wie der JHWH also."}
    KossonSILVER0403 = Kanon {"Kanon 15.  26", "Du bist leidend in Krankheit und leidend in tiefster Betrübtheit, und qualvoll sind oft die Stunden deiner Taten und Bemühungen, so du einsam bist als Mensch unter Menschen."}
    KossonSILVER0404 = Kanon {"Kanon 15.  27", "Nicht aber wirst du wankelmütig, und nicht lassest du brüchig werden dein Wort der Pflicht und Pflichterfüllung, und nicht weichet wahrlich ein Jota deines Willens und deiner Bemühungen zur Belehrung und zum Fortschritt und zur Missionserfüllung."}
    KossonSILVER0405 = Kanon {"Kanon 15.  28", "Also wird deiner Aufwendung beschieden sein Erfolg, gegeben in kleinem Masse während deiner Lebenszeit, und in grossem Masse nach deinem neuerlichen Vergehen.»"}

    //Kanon 16
    KossonSILVER0406 = Kanon {"Kanon 16.   1", "Im Namen der Schöpfung, der Weisen, der Gerechten."}
    KossonSILVER0407 = Kanon {"Kanon 16.   2", "Preis (Ehre) sei der Schöpfung, die da ist Freiheit."}
    KossonSILVER0408 = Kanon {"Kanon 16.   3", "Höre, nur jener Mensch, der da lebet in Freiheit, der allein lebet nach den Gesetzen und Geboten der Schöpfung, nur er kann wahrlich leben."}
    KossonSILVER0409 = Kanon {"Kanon 16.   4", "Und nur jener Mensch kann aufnehmen und sich öffnen der Wahrheit, der da wahrlich lebet in Freiheit und der da ist willig zu leben in Freiheit."}
    KossonSILVER0410 = Kanon {"Kanon 16.   5", "Es kann nur aufnehmen und zuhören jener Mensch, der da ist frei, und nur er kann sich öffnen der Wahrheit, weil er da frei ist im Willen so zu tun."}
    KossonSILVER0411 = Kanon {"Kanon 16.   6", "Tot im Leben sind aber alle jene, die da ihre Freiheit verschliessen und sie durch Unwahrheit und Ungerechtigkeit in Fesseln legen."}
    KossonSILVER0412 = Kanon {"Kanon 16.   7", "Sie machen sich zu Gefangenen ihrer selbst aus eigener Kraft und Falschheit."}
    KossonSILVER0413 = Kanon {"Kanon 16.   8", "Und sie sind lebende Tote, die da nicht erwecket werden können, so ihnen der eigene Sinn und Wille nicht danach stehet."}
    KossonSILVER0414 = Kanon {"Kanon 16.   9", "Und sie sind jene, die sich da wundern und unverstehend sind gegenüber jenen, die da sind erfolgreich in ihrer Freiheit."}
    KossonSILVER0415 = Kanon {"Kanon 16.  10", "Und es suchen jene, die sich selbst geknechtet haben, ihre Freiheit anderswo, als sie da zu finden ist im Menschen selbst."}
    KossonSILVER0416 = Kanon {"Kanon 16.  11", "Nicht ist der Mensch wissend und erkennend, dass die innere Freiheit ist die wahrliche Freiheit, so er in Ketten oder im Kerker liegen kann und doch in Freiheit lebet."}
    KossonSILVER0417 = Kanon {"Kanon 16.  12", "Es ist aber Freiheit im Innern des Menschen, wenn er gerecht lebet mit der Wahrheit und in Erfüllung der Gesetze und Gebote der Schöpfung."}
    KossonSILVER0418 = Kanon {"Kanon 16.  13", "Also ist da die wahrliche Freiheit, wenn der Mensch lebet in Freiheit in sich selbst, wenn ihm ist Wahrheit, Wissen, Liebe und Weisheit eigen, so er die Gesetze und Gebote der Schöpfung erfüllet in wahrlicher Freude der Pflicht."}
    KossonSILVER0419 = Kanon {"Kanon 16.  14", "Und es hat der Mensch frei zu sein in seinem Innern, so allein er also lebet in wahrheitlicher Freiheit."}
    KossonSILVER0420 = Kanon {"Kanon 16.  15", "Und so der Mensch ist frei in seinem Innern, sich ihm dann auch gestaltet die Freiheit im Ausseren, im Leben des Täglichen also."}
    KossonSILVER0421 = Kanon {"Kanon 16.  16", "Und ist der Mensch gerecht und in Freiheit in seinem Innern, dann lebet er ohne Anstoss und ohne Frevel und ohne Gesetzesbruch im Äusseren, so er weder nach innen noch nach aussen ungerecht ist und der Ungerechten einer wäre."}
    KossonSILVER0422 = Kanon {"Kanon 16.  17", "Und so er ist ein Gerechter, lebet er als Gerechter in seinem Innern und in seinem Äussern und in wahrlicher Freiheit ebenso."}
    KossonSILVER0423 = Kanon {"Kanon 16.  18", "Als Gerechter in Freiheit erreget er nicht Unmut und Unwillen im Innern oder Äussern, denn er ist nicht der Fehlbaren einer, die da wahrlich allein die Gerechten in Ungerechtigkeit befeinden und ihnen Übel wollen."}
    KossonSILVER0424 = Kanon {"Kanon 16.  19", "In Gemeinschaft lebet der Mensch mit seinesgleichen, doch in Freiheit lebet der Freie nur unter Freien, denn Ungerechte sind unfrei und befeinden die Freiheit."}
    KossonSILVER0425 = Kanon {"Kanon 16.  20", "Nicht vor der Schöpfung erregt der Gerechte Anstoss und Unmut, denn befeindet wird er nur von den Ungerechten, die da wahrlich gleichtun sollten dem Getier, das da lebet auf der Erde."}
    KossonSILVER0426 = Kanon {"Kanon 16.  21", "Kein Getier gibt es auf der Erde, keinen Vogel, der auf zwei Schwingen dahinflieget, die nicht in Gemeinschaften wären gleich dem Menschen."}
    KossonSILVER0427 = Kanon {"Kanon 16.  22", "Und es ist kein Getier auf Erden, das da wäre unfrei in sich selbst, und das da wäre eine ungerechte Lebensform also."}
    KossonSILVER0428 = Kanon {"Kanon 16.  23", "Jegliches Getier auf Erden lebet freiheitlich in seinem Innern, und es ist frei in seinem Äusseren also, so es lebet nach der gegebenen Ordnung der Gesetze und Gebote der Schöpfung."}
    KossonSILVER0429 = Kanon {"Kanon 16.  24", "Und ist der Mensch nicht viel mehr als alles Getier, dem er ist übergeordnet?"}
    KossonSILVER0430 = Kanon {"Kanon 16.  25", "Doch es lieget der Wille des Menschen in seiner freien Entscheidung, zu leben als Ungerechter oder als Gerechter, so also als Mensch der wahrlichen Freiheit oder als Mensch, der da ist Gefangener seiner selbst."}
    KossonSILVER0431 = Kanon {"Kanon 16.  26", "Wahrlich, jener Mensch, der da lebet in Unfreiheit seiner selbst, der wird fristen ein qualvolles Leben, und er wird siechen jenseits der Liebe und Freude."}
    KossonSILVER0432 = Kanon {"Kanon 16.  27", "Wahrheitliche Liebe und Freude ist nur eigen und zugeordnet dem Gerechten, der da wahrlich lebet in Freiheit im Innern und im Äussern."}
    KossonSILVER0433 = Kanon {"Kanon 16.  28", "Dem Freien allein wird Gerechtigkeit zuteil, und so auch wahrliche Liebe, Weisheit, Wissen und Wahrheit und Frieden, denn er lebet in Erfüllung und Befolgung der siebenheitlichen schöpferischen Ordnung."}

    //Kanon 17
    KossonSILVER0434 = Kanon {"Kanon 17.   1", "Im Namen der Schöpfung, der Weisen, der Gerechten."}
    KossonSILVER0435 = Kanon {"Kanon 17.   2", "Preis (Ehre) sei der Schöpfung, der Allumfassenden."}
    KossonSILVER0436 = Kanon {"Kanon 17.   3", "«Siehe», spricht der JHWH, «nichts ist im Buch der Wahrheit ausgelassen («Geisteslehre» und «Kelch der Wahrheit» usw.) und nichts bleibet dem Menschen verborgen, wenn er sich willig bemühet um die Kenntnis der Schrift der Wahrheit des Wortes."}
    KossonSILVER0437 = Kanon {"Kanon 17.   4", "Allumfassend ist das Wort der Wahrheit und die Lehre des Geistes, so der Mensch das Wort erfassen kann und das Leben wahrlich leben möge."}
    KossonSILVER0438 = Kanon {"Kanon 17.   5", "Gegeben ist dem Menschen durch die Schöpfung die Freiheit des Willens, und gegeben ist ihm die Kraft und die Macht, so er alles nutze gerecht zu seinem Fortschritt und zum Zielstreben der Evolution."}
    KossonSILVER0439 = Kanon {"Kanon 17.   6", "Und es sind ihm auch gegeben die Richtlinien und die Gesetze und Gebote der Schöpfung, so er habe die nutzvollen Wegweiser in Unfehlbarkeit zur Erfüllung des Lebens."}
    KossonSILVER0440 = Kanon {"Kanon 17.   7", "Der Mensch aber möge gemäss dem Schöpfungswillen selbst walten und fortschreitend sein nach eigenem Ermessen, so ihm kein Zwang auferleget sei durch Gesetze und Gebote der Menschen in dieser Sache, und so er frei sei in wahrlicher Freiheit."}
    KossonSILVER0441 = Kanon {"Kanon 17.   8", "Er möge aber befolgen die Richtlinien und die schöpferisch-natürlichen Gesetze und Gebote, so er lebend sei und fortschreitend in gerechter und wahrer Form."}
    KossonSILVER0442 = Kanon {"Kanon 17.   9", "So der Mensch aber nicht handelt danach, dann folget er den Wegen der Ungerechten, denen er eigene Strafe zollet alsbald für sein Handeln im Fehlbaren."}
    KossonSILVER0443 = Kanon {"Kanon 17.  10", "Und wenn gesaget ist also, dass dem Menschen frei sei sein Wille und seine Kraft und Macht und das Denken und Handeln zur freien Entscheidung und Nutzung, dann ist dies gegeben im Zeichen der freien Wahl der richtigen Befolgung der siebenheitlichen Ordnung der Schöpfung."}
    KossonSILVER0444 = Kanon {"Kanon 17.  11", "Also besaget die Lehre, dass gegeben sind die schöpferischen siebenheitlichen Gesetze und Richtlinien, die da sollen befolget sein durch den Menschen und das Leben im gesamten."}
    KossonSILVER0445 = Kanon {"Kanon 17.  12", "Doch die Wege der Befolgung der Richtlinien und Gesetze und Gebote der Schöpfung seien dem Menschen freigestellet, denn der Wege der Befolgung sind unzählbare und so zahlreich, wie Menschen sind existent im Himmel (Universum)."}
    KossonSILVER0446 = Kanon {"Kanon 17.  13", "Also ist nicht gegeben eine Ordnung dessen, welcher Weg da soll beschritten werden zur Befolgung und Erfüllung der Richtlinien und Gesetze und Gebote, sondern es ist nur gegeben, dass begangen werden wahrheitliche und gesetzmässige Wege, die da führen zur Erfüllung der Gesetze und Gebote."}
    KossonSILVER0447 = Kanon {"Kanon 17.  14", "Es ist der Mensch also frei und ohne Zwang oder Gewalt in der Wahl seiner Wege, seines Denkens und seines Handelns, wenn er nur fortschreitet und lebet in Befolgung der Erfüllung der siebenheitlichen Schöpfungsordnung."}
    KossonSILVER0448 = Kanon {"Kanon 17.  15", "Niemals ist auferleget dem Menschen Gewalt und Zwang durch die Schöpfung oder durch ihre Gesetze und Gebote oder durch ihre siebenheitliche Ordnung."}
    KossonSILVER0449 = Kanon {"Kanon 17.  16", "Und lebet der Mensch gerecht und nach der Ordnung der Schöpfung, dann erfüllet sich ihm ein Leben in Liebe und Frieden, und so in Wahrheit und Wissen und Weisheit also."}
    KossonSILVER0450 = Kanon {"Kanon 17.  17", "Wandelt aber der Mensch auf Wegen des Ungerechten, und verstosset er gegen die Ordnung der Schöpfung zur Erfüllung von Eigensucht, Machtgier, Vorteil und Profitgier und gleicher Dinge mehr, dann zeuget er sich Schaden und Schuld in eigener Kraft, denen er verfallet also."}
    KossonSILVER0451 = Kanon {"Kanon 17.  18", "Wege des Ungerechten sind Verstoss gegen die Gesetze und Gebote der Schöpfung und gegen die siebenheitliche Ordnung also."}
    KossonSILVER0452 = Kanon {"Kanon 17.  19", "Und so der Mensch Verstoss übet wider die Schöpfungsordnung, so verstosset er gegen das Allumfassende und zeuget sich Schuld und Schaden."}
    KossonSILVER0453 = Kanon {"Kanon 17.  20", "Begehet der Mensch Verstoss gegen das Höchste, dann erhebet er sich in Hochmut, und er lästert der Wahrheit und Existenz der Schöpfung."}
    KossonSILVER0454 = Kanon {"Kanon 17.  21", "Hochmut führet zu Last und Schuld, und die Bürde des Hochmuts drücket bis zum Fall, so auch gesaget ist, dass Hochmut kommet vor dem Fall.»"}

    //Kanon 18
    KossonSILVER0455 = Kanon {"Kanon 18.   1", "Im Namen der Schöpfung, der Weisen, der Gerechten."}
    KossonSILVER0456 = Kanon {"Kanon 18.   2", "Preis (Ehre) sei der Schöpfung, die da ist Tröstung."}
    KossonSILVER0457 = Kanon {"Kanon 18.   3", "Und Trost findet der Mensch in der Schöpfung, so in ihren Gesetzen und Geboten und in der siebenheitlichen Ordnung also."}
    KossonSILVER0458 = Kanon {"Kanon 18.   4", "Es ist des Lebens Trost in aller Grösse, wenn es gelebet wird in siebenheitlicher Ordnung."}
    KossonSILVER0459 = Kanon {"Kanon 18.   5", "So aber ergehet das Wort des JHWH an die Menschen, das da besaget seit alters her und abermals zu neuer Zeit, dass da ist die Wahrheit in der Schöpfung."}
    KossonSILVER0460 = Kanon {"Kanon 18.   6", "Und es sind selig alle, die da geistig reich sind und die da die Wahrheit erkennen und befolgen, denn ihrer ist das Leben."}
    KossonSILVER0461 = Kanon {"Kanon 18.   7", "Und es sind selig alle, die da ertragen Leid und gerechte Mühsal, denn sie erkennen daraus die Wahrheit, und in ihr finden sie Trost."}
    KossonSILVER0462 = Kanon {"Kanon 18.   8", "Und es sind selig alle, die da geistig ausgeglichen sind, denn sie sind wahrheitlich Wissende und leben in Ehrfurcht (Ehrung, Ehrwürdigkeit, Ehrerbietung, Ehrwürdigung)."}
    KossonSILVER0463 = Kanon {"Kanon 18.   9", "Und es sind selig alle, die da hungern und dürsten nach Wahrheit und Wissen, denn sie werden Wahrheit und Wissen finden und durch sie gesättiget werden."}
    KossonSILVER0464 = Kanon {"Kanon 18.  10", "Und es sind selig alle, die den Gesetzen der Natur folgsam sind und ihnen nachleben, denn sie leben wahrheitlich nach dem Schöpfungsplan."}
    KossonSILVER0465 = Kanon {"Kanon 18.  11", "Und selig sind alle, die gerecht und reinen Gewissens sind, denn ihrer harret weder eigenes Strafgericht noch Angst oder Furcht."}
    KossonSILVER0466 = Kanon {"Kanon 18.  12", "Und selig sind alle, die wahrheitlich wissend sind um die Schöpfung, denn ihnen ist es fern, einer irren Lehre und Götzen und Göttern zu frönen."}
    KossonSILVER0467 = Kanon {"Kanon 18.  13", "Und selig sind alle, die da gerechte Gesetzestreue sind, denn ihnen unterordnet sich die Natur."}
    KossonSILVER0468 = Kanon {"Kanon 18.  14", "Und selig sind alle, die da gerechte Schöpfungstreue sind, denn ihnen werden nutzbar die Kräfte der Schöpfung."}
    KossonSILVER0469 = Kanon {"Kanon 18.  15", "Und selig sind alle, die da fallen in Schmähung und Verfolgung um der Wahrheit willen, denn sie bestehen durch die Kraft der Wahrheit der Schöpfung."}
    KossonSILVER0470 = Kanon {"Kanon 18.  16", "Und selig sind alle, die da Leid und Mühsal tragen um der Wahrheit willen, denn sie sind beständig, wissend und unbeugsam in der Wahrheit, so ihnen daraus Trost entstehet."}
    KossonSILVER0471 = Kanon {"Kanon 18.  17", "Und es belehret der JHWH in Deutung der Gesetze und Gebote der Schöpfung, dass der Mensch soll sein wahrheitlich und ehrfürchtig im Wissen um die Schöpfung, und liebentbietend und ehrentbietend also."}
    KossonSILVER0472 = Kanon {"Kanon 18.  18", "Es lebe der Mensch einzig in Gerechtigkeit und in der Wahrheit der Schöpfung, so ihm werde das Leben zum SEIN (Schöpfungsleben) durch die Evolution."}
    KossonSILVER0473 = Kanon {"Kanon 18.  19", "Es sei der Mensch fröhlich und getrost in Erfüllung der siebenheitlichen schöpferischen Ordnung, und es wird ihm dies gelohnet sein durch das Leben und das Wiederleben."}
    KossonSILVER0474 = Kanon {"Kanon 18.  20", "Der Mensch öffne sein Ohr gegenüber seinem Propheten, der da beständig ist und die Wahrheit bringet, obwohl er wird verleumdet und verfolget, so er auch verleumdet und verfolget war ehedem."}
    KossonSILVER0475 = Kanon {"Kanon 18.  21", "Seid gerecht in allen Dingen, und höret nicht der Schmähungen, und achtet nicht der Verfolgungen wider den Propheten und euch, so die Schmähungen und Verfolgungen getan sind um der Wahrheit willen."}
    KossonSILVER0476 = Kanon {"Kanon 18.  22", "Seid beständig, ihr Menschen, denn die Wahrheit muss erkämpfet sein wider die Ungerechtigkeit des Bösen, das da ist Widerwille, Hass und Widerspenstigkeit des Menschen gegen die Wahrheit und Existenz er Schöpfung."}
    KossonSILVER0477 = Kanon {"Kanon 18.  23", "Die Wahrheit ist beständiges Wissen, und sie ist gewürzet mit Liebe und Weisheit, so wie auch der Mensch würzet sein Mahl, dass es ihm wird geniessbar."}
    KossonSILVER0478 = Kanon {"Kanon 18.  24", "Und wie Liebe und Weisheit sind die Würze der Wahrheit und des Wissens, so sind die Menschen die Würze der Erde."}
    KossonSILVER0479 = Kanon {"Kanon 18.  25", "Und also seid ihr die Würze; ihr seid das Salz der Erde."}
    KossonSILVER0480 = Kanon {"Kanon 18.  26", "Und ihr Menschen, ihr würzet und salzet die Erde je nachdem, wie eure Lebensführung dies bestimmet."}
    KossonSILVER0481 = Kanon {"Kanon 18.  27", "So ihr das Salz der Erde seid, und ihr salzet die Erde in gutem Mass, dann habet ihr ein Gleichmass geschaffen."}
    KossonSILVER0482 = Kanon {"Kanon 18.  28", "Und so ihr das Salz der Erde seid, und ihr salzet die Erde im Übermass, dann wird sie durstig, rissig, brüchig und trocknet aus."}
    KossonSILVER0483 = Kanon {"Kanon 18.  29", "Und so ihr das Salz der Erde seid, und das Salz wird kraftlos, womit salzet ihr dann die Erde, denn kraftloses Salz ist hinfort zu nichts mehr nutze, so es hinweggeschüttet und zertreten werde."}
    KossonSILVER0484 = Kanon {"Kanon 18.  30", "Und so ihr das Salz der Erde seid, so seid ihr dies im Sinne eurer Lebensführung, denn so, wie ihr euer Leben gestaltet, so gestaltet ihr auch die Erde."}
    KossonSILVER0485 = Kanon {"Kanon 18.  31", "Gestaltet ihr also euer Leben in Liebe und Freude und in Erfüllung der Wahrheit der Schöpfung, so würzet ihr die Erde im Gleichmass desselben, und es wird sein die Erde für euch ein Paradies."}
    KossonSILVER0486 = Kanon {"Kanon 18.  32", "Gestaltet ihr aber euer Leben in Ungerechtigkeit, in Lieblosigkeit und Freudlosigkeit, in Eigensucht und Machtgier, in Vorteilheischung und Profitgier und in Sucht und Laster, dann würzet ihr die Erde mit Hass und Tod, mit Krieg und Krankheit, und mit Siechtum, Unfrieden, Zerstörung und Vernichtung also."}
    KossonSILVER0487 = Kanon {"Kanon 18.  33", "Gestaltet ihr aber euer Leben in erfüllender Gerechtigkeit und in Wahrheit der Gesetze und Gebote der Schöpfung, dann wird euch sein Glückseligkeit, so es euch dünket, als wären alle Herrlichkeiten der Himmel (höherer Ebenen) auf Erden, so euch schöpferischer Trost ist in unendlicher Folge."}
    KossonSILVER0488 = Kanon {"Kanon 18.  34", "Und wahrlich, ihr sollet das Salz der Erde sein, und die Erde sollet ihr salzen mit wahrlichem Leben."}
    KossonSILVER0489 = Kanon {"Kanon 18.  35", "Denket, dass ihr das Licht eurer Welt seid und dass der Schein des Lichtes sich derart ins All verbreitet, wie ihr dieses Licht gestaltet."}
    KossonSILVER0490 = Kanon {"Kanon 18.  36", "Erzeuget ihr das Unlicht der Dunkelheit, so finstert es dunkel in den Raum."}
    KossonSILVER0491 = Kanon {"Kanon 18.  37", "Und erzeuget ihr das Licht der Helligkeit, dann hellet es die Dunkelheit und strahlet weit durch den Raum, zeugend davon, dass wahrliches Leben blühet."}
    KossonSILVER0492 = Kanon {"Kanon 18.  38", "Siehe, Mensch der Erde, ein jegliches Leben zeuget von seiner Art und Weise, so im Guten wie im Bösen."}
    KossonSILVER0493 = Kanon {"Kanon 18.  39", "Und keine Art und Weise eines jeglichen Lebens kann verborgen bleiben, wie auch nicht eine Stadt verborgen bleiben kann, die auf einem Berge lieget."}
    KossonSILVER0494 = Kanon {"Kanon 18.  40", "Niemals entzünde ein Licht und stelle es unter einen Scheffel, so dunkles Licht in den Raum hineinfinstert."}
    KossonSILVER0495 = Kanon {"Kanon 18.  41", "Entzünde stets ein Licht, das getan wird in einen Leuchter, so es leuchten möge allen, die da des strahlenden Lichtes bedürfen."}
    KossonSILVER0496 = Kanon {"Kanon 18.  42", "Und das Licht soll nicht von Dunkelheit sein, sondern strahlend, leuchtend und hell."}
    KossonSILVER0497 = Kanon {"Kanon 18.  43", "Es soll so der Mensch rechtschaffen sein und gerecht und den Gesetzen und Geboten der Schöpfung erfüllend folgen, so er ausstrahlet das Licht und die Schwingung der Wahrheit."}
    KossonSILVER0498 = Kanon {"Kanon 18.  44", "Und es soll euer Licht leuchten vor allen Menschen, und es soll eure Schwingung liebbringend und friedbringend sein für alle Menschen, auf dass sie eure Werke der Wahrheit sehen und die Wahrheit und das Wissen eurer Liebe und Weisheit erkennen und verstehen."}
    KossonSILVER0499 = Kanon {"Kanon 18.  45", "Und ihr sollet so leben, denken und handeln in alle Zeit des Kommenden, so gegenwärtig und in alle Zukunft."}
    KossonSILVER0500 = Kanon {"Kanon 18.  46", "Und ihr sollet nicht wähnen, dass das Wort der Wahrheit und euch neuerlich ein Prophet gegeben sei, dass aufgelöset würden die Gesetze und Gebote der Schöpfung und die Lehren des JHWH und der Propheten, die da waren vor eurer Zeit."}
    KossonSILVER0501 = Kanon {"Kanon 18.  47", "Es ist der neue Prophet nicht gegeben zur Auflösung der seit alters her gegebenen Lehren, sondern zu erfüllen ein andermal die Pflicht der Offenbarung des Wissens, der Wahrheit, der Liebe und der Weisheit."}
    KossonSILVER0502 = Kanon {"Kanon 18.  48", "Denn wahrlich, es ist die Pflicht des JHWH und des Propheten, Trost zu geben allen jenen, die da willig und hoffend sind, dass sie die Wahrheit finden, so ein jeglicher soll finden, der da suchet."}
    KossonSILVER0503 = Kanon {"Kanon 18.  49", "Und wahrlich, die Wahrheit ist findbar für einen jeglichen, der da in Ehrlichkeit suchet und sich bemühet."}
    KossonSILVER0504 = Kanon {"Kanon 18.  50", "Und wer da findet die Wahrheit in wahrlicher Form, der erlanget die Erkenntnis und das Wissen, dass Liebe, Wahrheit, Wissen und Weisheit gegeben sind durch die Existenz der Schöpfung, die da ist unvergänglich und unveränderbar in der Wahrheit und in allem Sonstigen, das ihr ist eigen."}
    KossonSILVER0505 = Kanon {"Kanon 18.  51", "Und siehe: Bis dass die Himmel (Universum, Ebenen und Dimensionen) und die Erden vergehen, wird nicht vergehen der kleinste Buchstabe noch ein Tüpfelchen von der Wahrheit und von den Gesetzen und Geboten der Schöpfung, so aber auch nicht von den Gesetzen und Geboten der Natur und von der siebenheitlichen Ordnung, bis dass sich erfüllet die Zeit der Zeit."}
    KossonSILVER0506 = Kanon {"Kanon 18.  52", "Wer unter den Menschen aber einer ist, der da versuchet, auch nur eines der kleinsten Gesetze oder Gebote der Schöpfung aufzulösen, und der da bringet den Menschen eine irrige Lehre, der wird der Kleinste heissen und die Schuld des eigenen Gerichtes auf sich laden."}
    KossonSILVER0507 = Kanon {"Kanon 18.  53", "Wer aber die Lehre der Wahrheit und die Gesetze und Gebote der Schöpfung wahrheitlich belehret unter den Menschen, der wird der Grösste heissen und den Dank der Menschen und des Geistes ernten."}
    KossonSILVER0508 = Kanon {"Kanon 18.  54", "Seid ihr nicht besser und wissender als die Ungerechten, dann werdet ihr nicht erhalten den Dank der Menschen, und nicht den Dank des Geistes (Bewusstseins) und nicht den Dank des Lebens."}
    KossonSILVER0509 = Kanon {"Kanon 18.  55", "Ihr wisset, dass gesaget ist: Du sollest nicht töten; wer da aber tötet, der machet sich der selbstverschuldeten Strafe und des eigenen Gerichtes schuldig."}
    KossonSILVER0510 = Kanon {"Kanon 18.  56", "Daher sollet ihr üben Gerechtigkeit nach dem Naturgesetz der Schöpfung, so ihr nicht tötet aus Rache, aus Habsucht, aus Hass oder nach einem Urteil oder aus anderen Gründen der Nichtigkeiten."}
    KossonSILVER0511 = Kanon {"Kanon 18.  57", "Findet eine jegliche Klarheit in den Gesetzen und Geboten der Natur, so ihr Beurteilung findet in der Logik, die da euch dann ist Trost und Gerechtigkeit."}
    KossonSILVER0512 = Kanon {"Kanon 18.  58", "Und schuldbar werden da alle, die da töten, wenn sie nicht in gegebener Notwehr handeln."}
    KossonSILVER0513 = Kanon {"Kanon 18.  59", "Und schuldbar werden da alle, die nach menschlich erdachten Gesetzesurteilen vollstreckend handeln und töten."}
    KossonSILVER0514 = Kanon {"Kanon 18.  60", "Und schuldbar werden da alle, die da töten aus Freude oder aus Hass, aus Rache, aus Gewinnsucht, aus Eifer und anderen Emotionen."}
    KossonSILVER0515 = Kanon {"Kanon 18.  61", "Und schuldbar werden da alle, die da töten als Kriegshandelnde in jeglicher Form, wenn sie sind die angreifende Kraft."}
    KossonSILVER0516 = Kanon {"Kanon 18.  62", "Und schuldbar werden da alle, die da töten in jeglicher Ausartung."}
    KossonSILVER0517 = Kanon {"Kanon 18.  63", "Wahrlich, allein die Gerechtigkeit der Gesetze und Gebote der Schöpfung erheben eine Beurteilung in die Logik, die da erkenntlich ist in den Naturgesetzen."}
    KossonSILVER0518 = Kanon {"Kanon 18.  64", "Stets sollet ihr gerecht sein, so im Nehmen wie im Geben."}
    KossonSILVER0519 = Kanon {"Kanon 18.  65", "So sollet ihr auch nicht willfährig sein euren Widersachern, wenn ihr im Recht seid, denn das Recht muss des Rechtes bleiben."}
    KossonSILVER0520 = Kanon {"Kanon 18.  66", "Weder Angst noch Furcht, noch Niedertracht, noch Vorteil oder Nachteil halte euch vor dem Recht zurück, so ihr immer sollet nach Gerechtigkeit streben und nur dieser Geltung verschaffen."}
    KossonSILVER0521 = Kanon {"Kanon 18.  67", "So aber ein Richter in einer eurer Sachen rechtet und richtet und zu falschem Urteil kommet, so rechtet weiter bis zum Sieg der gerechten Sache."}
    KossonSILVER0522 = Kanon {"Kanon 18.  68", "Doch wahrlich sei gesagt: Ihr werdet stetig nur dann Gerechtigkeit erlangen, wenn ihr selbst Gerechtigkeit findet in euch selbst, so ihr gerecht seid in eigener Kraft und so ihr die Gerechtigkeit eurem Nächsten könnet machen verständlich."}
    KossonSILVER0523 = Kanon {"Kanon 18.  69", "Es findet der Gerechte Gerechtigkeit, wenn er ist kämpfend um die Gerechtigkeit mit den Mitteln der Gerechtigkeit."}
    KossonSILVER0524 = Kanon {"Kanon 18.  70", "Und es findet der Gerechte Trost in der Gerechtigkeit, wenn er wahrlich ist gerecht in sich selbst und also lebet in Gerechtigkeit."}

    //Kanon 19
    KossonSILVER0525 = Kanon {"Kanon 19.   1", "Im Namen der Schöpfung, der Weisen, der Gerechten."}
    KossonSILVER0526 = Kanon {"Kanon 19.   2", "Preis (Ehre) sei der Schöpfung, die da ist Ganzheit."}
    KossonSILVER0527 = Kanon {"Kanon 19.   3", "Und es sprichet der JHWH, dass da die Schöpfung ist Ganzheit, und so sind Ganzheit ihre Gesetze und Gebote und die siebenheitliche Ordnung also."}
    KossonSILVER0528 = Kanon {"Kanon 19.   4", "Ganzheit aber ist das, das da ist genennet das Allumfassende in Vollständigkeit."}
    KossonSILVER0529 = Kanon {"Kanon 19.   5", "In Vollständigkeit ist die Schöpfung Ganzheit,und so ihre Gesetze und Gebote und die siebenheitliche Ordnung also."}
    KossonSILVER0530 = Kanon {"Kanon 19.   6", "Ganzheit ist verstehend zu erkennen als Vollständigkeit, als Allumfassendes, als Einheit und als Gesamtheit, und es ist die Ganzheit auch die Gleichform, die Eintracht, der Einklang und die Verschmolzenheit."}
    KossonSILVER0531 = Kanon {"Kanon 19.   7", "Also ist die Schöpfung Ganzheit, und es sind Ganzheit ihre Gesetze und Gebote und ihre siebenheitliche Ordnung, und so der Mensch danach lebet, lebet er in der Ganzheit der Richtlinien und Erfüllung des Lebens."}
    KossonSILVER0532 = Kanon {"Kanon 19.   8", "Es ist auch das Wort der Wahrheit Ganzheit in Vollständigkeit, denn es offenbaret in wahrheitlicher Form die Auslegung, Ausdeutung und Erklärung der Gesetze und Gebote und die siebenheitliche Ordnung der Schöpfung."}
    KossonSILVER0533 = Kanon {"Kanon 19.   9", "Und was da ist Ganzheit und Vollständigkeit, das bleibet unveränderbar für alle Zeit."}
    KossonSILVER0534 = Kanon {"Kanon 19.  10", "Also ist auch unveränderbar alle Ganzheit der Schöpfung."}
    KossonSILVER0535 = Kanon {"Kanon 19.  11", "Und es ist unveränderbar die Ganzheit der Gesetze und Gebote der Schöpfung."}
    KossonSILVER0536 = Kanon {"Kanon 19.  12", "Und es ist unveränderbar die Ganzheit der siebenheitlichen Ordnung der Schöpfung."}
    KossonSILVER0537 = Kanon {"Kanon 19.  13", "Und es ist unveränderbar die Ganzheit der Ganzheit, die Wahrheit, die Liebe, das Wissen und die Weisheit in ihren wahrheitlichen Formen."}
    KossonSILVER0538 = Kanon {"Kanon 19.  14", "Also ist dadurch auch gegeben, dass da ist unveränderbar das Wort (Inhalt) der Wahrheit, das da ist gegeben seit alters her in gleichbleibendem Sinn und Wert."}
    KossonSILVER0539 = Kanon {"Kanon 19.  15", "Und es beruhet das Buch des Wortes der Wahrheit einzig und allein auf der wahrheitlichen Auslegung, Erkennung und Erklärung und Ausdeutung der Gesetze und Gebote und der siebenheitlichen Ordnung der Schöpfung."}
    KossonSILVER0540 = Kanon {"Kanon 19.  16", "Und es sind die Erkennung, die Auslegung, die Ausdeutung und die Erklärung gegeben in reinster und wahrheitlicher Form der wahrlichen Kenntnis und Befolgung der Gesetze und Gebote und der siebenheitlichen Ordnung der Schöpfung."}
    KossonSILVER0541 = Kanon {"Kanon 19.  17", "Und es ist erkennet und erfahren, dass nichts ist veränderbar darin und dass die Ganzheit der Wahrheit ebenso ist unveränderbar wie auch die Schöpfung und ihre Gesetze und Gebote und die siebenheitliche Ordnung."}
    KossonSILVER0542 = Kanon {"Kanon 19.  18", "Und es sind die Erkennung, die Auslegung und die Ausdeutung der Ganzheit der Schöpfung und ihre Gesetze und Gebote und die siebenheitliche Ordnung erfasset worden von Menschen in Weisheit, Wissen,Liebe und Wahrheit, und ergründet, festgehalten und aufgezeichnet für den Menschen, so er erkenne und erfülle das Ziel seines Lebens zur Evolution zum SEIN （Schöpfungsleben)."}
    KossonSILVER0543 = Kanon {"Kanon 19.  19", "Und es wurde all das Wissen der Ganzheit ganzheitlich festgehalten und niedergeschrieben in dem Buch, das da ist genennet seit alters her das Buch des Wortes der Wahrheit, wie es also ward gegeben dem Erdenmenschen ehedem."}
    KossonSILVER0544 = Kanon {"Kanon 19.  20", "Und also sind gegeben seit alters her den Erdenmenschen andere ausführende Niederschriften der Wahrheit der Ganzheit, worin sind aufgeteilet und erkläret in verstehender und gesonderter Form die Gesetze und Gebote und die siebenheitliche Ordnung der Schöpfung."}
    KossonSILVER0545 = Kanon {"Kanon 19.  21", "Also sind auch gegeben die gesonderten und verständlichen Erklärungen, Auslegungen, Ausdeutungen und Erkenntnisse der Gesetze und Gebote und der siebenheitlichen Ordnung in ausführlicher und getrennter Form, so also entnommen aus dem Buche des Wortes der Wahrheit,und festgehalten und niedergeschrieben in der Schrift, die da genennet ist die Lehre des Geistes."}
    KossonSILVER0546 = Kanon {"Kanon 19.  22", "Es ist erkenntlich, dass die Ganzheit der Schöpfung ist vollständig, und so auch ihre Gesetze und Gebote und also ihre siebenheitliche Ordnung, die da sind gleichbleibend und unveränderlich allzeitlich."}
    KossonSILVER0547 = Kanon {"Kanon 19.  23", "Und es ist kein Zweifel an der wahrheitlichen Wahrheit der Ganzheit der Schöpfung, und an ihren Gesetzen und Geboten und der siebenheitlichen Ordnung also."}
    KossonSILVER0548 = Kanon {"Kanon 19.  24", "Und nicht und niemals ist gegeben eine Möglichkeit, die da zulassen würde eine Handhabe zur Auslegung, Erklärung und Ausdeutung der Wahrheit und der Ganzheit der Schöpfung und ihrer Gesetze und Gebote und der siebenheitlichen Ordnung in zwei verschiedene Formen und Erklärungen."}
    KossonSILVER0549 = Kanon {"Kanon 19.  25", "Nicht und niemals können gegeben sein zwei verschiedene Deutungen der Wahrheit oder zwei Erklärungen der Wahrheit also, denn die Wahrheit ist wahrheitlich und einzig, und so erklärbar auch nur in einem Wert und in einem Sinn."}
    KossonSILVER0550 = Kanon {"Kanon 19.  26", "Also kann die allumfassende und vollständige Ganzheit der Schöpfung und ihre Gesetze und Gebote und ihre siebenheitliche Ordnung nur aufweisen eine Erklärung, eine Auslegung und eine Deutung, denn da keine Zweifel gegeben sind in der Wahrheit, so ist nur eine Möglichkeit gegeben als einzige und von Richtigkeit in Wahrheit."}
    KossonSILVER0551 = Kanon {"Kanon 19.  27", "Und wie die Schöpfung Einheit ist, sind auch einheitlich ihre Gesetze und Gebote und ihre siebenheitliche Ordnung, so aber auch die Liebe, die Wahrheit, das Wissen und die Weisheit in Folgerichtigkeit."}
    KossonSILVER0552 = Kanon {"Kanon 19.  28", "Niemals können so sein zwei unterschiedliche Anschauungen über die gleiche Sache, und so also auch nicht zwei verschiedene Gruppierungen von Sinnesrichtungen unter den Menschen, wenn es da gehet um die Wahrheit."}
    KossonSILVER0553 = Kanon {"Kanon 19.  29", "Es kann erkennet und erfasset werden die Wahrheit nur als die eine und die wahrliche Wahrheit,so also als nur eine Anschauung und als nur eine Ansicht, als nur eine Meinung, als nur eine Auffassung, als nur eine Sinnesrichtung, als nur eine Gewissheit, als nur ein Ermessen, als nur ein Gesichtspunkt, als nur ein Gutdünken, als nur ein Standpunkt, als nur eine Vorstellung, als nur eine Stellungnahme, als nur eine Denkweise, als nur eine Gesinnung, als nur eine Warte und als nur ein Erachten."}
    KossonSILVER0554 = Kanon {"Kanon 19.  30", "So nur eine Schöpfung ist für den gesamten Himmel (Universum) und alles Heer des Lebens auf allen Welten, so ist auch nur eine Wahrheit in unveränderlicher Form gegeben, wobei Gleiches auch geltend ist für die Gesetze und Gebote und die siebenheitliche Ordnung der Schöpfung."}
    KossonSILVER0555 = Kanon {"Kanon 19.  31", "Und so der Mensch Religionen machet und sich Götter und Götzen berufet, so handelt er falsch und wahrheitsfremd und wider die Wahrheit der Ganzheit, und er erschaffet und frönet verschiedenartigen falschen Sinnesrichtungen, die da aufgebauet sind auf Irrlehren von Falschheiten und die da sind fern der Wahrheit also."}
    KossonSILVER0556 = Kanon {"Kanon 19.  32", "So wahrlich nur eine Schöpfung ist gegeben mit ihren Gesetzen und Geboten und der siebenheitlichen Ordnung im Himmel (Universum), so ist gegeben auch nur eine einzige Wahrheit als Sinnesrichtung, die da ist gleichbleibend und unveränderbar für alle Zeit und für den ganzen Himmel (Universum) und für alle Heere des Lebens also."}
    KossonSILVER0557 = Kanon {"Kanon 19.  33", "Die Schöpfung ist die Schöpfung,und die Wahrheit ist die einzige Wahrheit, so gegeben ist auch nur eine einzige Sinnesrichtung, und niemals deren zwei oder viele:"}
    KossonSILVER0558 = Kanon {"Kanon 19.  34", "Denn es ist die Schöpfung Ganzheit, und Ganzheit sind ihre Gesetze und Gebote und ihre siebenheitliche Ordnung."}
    KossonSILVER0559 = Kanon {"Kanon 19.  35", "Und was Ganzheit ist, wahrlich,das hat allzeitlich nur eine Wahrheit und eine Sinnesrichtung, denn die Ganzheit in ihrer Ganzheit kann niemals geteilet oder unterteilet werden in verschiedene Wahrheitsauslegungen und auch nicht in verschiedene Warten und Anschauungen, denn eine Wahrheit birget wahrlich in sich nur eine Wahrheit."}

    //Kanon 20
    KossonSILVER0560 = Kanon {"Kanon 20.   1", "Im Namen der Schöpfung, der Weisen, der Gerechten."}
    KossonSILVER0561 = Kanon {"Kanon 20.   2", "Preis (Ehre) sei der Schöpfung, die da ist Belehrung."}
    KossonSILVER0562 = Kanon {"Kanon 20.   3", "«Höre», sprichet der JHWH zum Propheten, «du bist gesendet unter die Menschen, so sie durch dich belehret werden über die Wahrheit der Ganzheit der Schöpfung:"}
    KossonSILVER0563 = Kanon {"Kanon 20.   4", "Die da ist die Kraft allen Lebens, und der weit untergeordnet ist als Mensch der JHWH, wie auch der Prophet."}
    KossonSILVER0564 = Kanon {"Kanon 20.   5", "Schau jedoch, wie die Menschen lügen wider sich selbst."}
    KossonSILVER0565 = Kanon {"Kanon 20.   6", "Und schau, wie sich alles gegen sie stellet und wie sie im Stiche lassen, was sie sich selbst ersannen an Glauben, Religionen, Sekten, Götzen und Göttern, die da wahrlich nur sind allesamt menschliche irre Machenschaften des Ungerechten."}
    KossonSILVER0566 = Kanon {"Kanon 20.   7", "Unter ihnen sind manche, die dir Gehör schenken, doch auf ihren Sinnen liegen selbstgefertigte Hüllen des Unverstehens und des Unbegreifens, und ihre Ohren sind voller Taubheit, so sie nicht hören, und ihre sehenden Augen sind verschlossen, so sie nicht sehen."}
    KossonSILVER0567 = Kanon {"Kanon 20.   8", "Selbst wenn sie wahrheitliche Zeichen sehen, so leugnen sie diese, und sie verstehen sie nicht."}
    KossonSILVER0568 = Kanon {"Kanon 20.   9", "So kommen sie zu dir, um mit dir in Ungerechtigkeit zu streiten, und sie verfolgen und beschimpfen dich."}
    KossonSILVER0569 = Kanon {"Kanon 20.  10", "Und es behaupten die Unwissenden und die Ungerechten, dass deine Worte der Wahrheit und deine Beweisführung nur Fabeln wären und wilde Phantasie und Betrug."}
    KossonSILVER0570 = Kanon {"Kanon 20.  11", "Doch stehen sie dereinst am Tor der Not und der Drangsal oder am Tor des Todes, dann wünschen sie, dass sie zurückgebracht würden, und sie versprechen, dass sie dann nicht mehr die Zeichen der Wahrheit und die Worte der Wahrheit der Lüge bezichtigen würden."}
    KossonSILVER0571 = Kanon {"Kanon 20.  12", "Ihr Sprechen ist jedoch nur Schein und Lüge, denn ganz gewiss würden sie bald zu ihrem falschen Tun und Denken zurückkehren, und so auch zu ihren Göttern und Götzen."}
    KossonSILVER0572 = Kanon {"Kanon 20.  13", "Und sie würden denken, handeln und sprechen wie ehedem, dass es kein anderes als das materielle irdische Leben gäbe und dass sie nach dem Tode nicht als neue Persönlichkeit erwecket würden zum Wiederleben in einem neuen irdischen Menschenkörper."}
    KossonSILVER0573 = Kanon {"Kanon 20.  14", "Spreche so zu diesen Unwissenden: «lhr seid die wahrlichen Verlierer, die ihr die Existenz und die Wahrheit der Schöpfung und ihre Gesetze und Gebote und ihre siebenheitliche Ordnung leugnet."}
    KossonSILVER0574 = Kanon {"Kanon 20.  15", "Bedenket aber: Wenn die Stunde der Rechenschaft über euch kommet, dann wehe euch Ungerechten, dass ihr die Wahrheit verleugnet und vernachlässiget habet."}
    KossonSILVER0575 = Kanon {"Kanon 20.  16", "Und wahrlich, wisset, dass ihr eure Last des ungerechten Lebens auf euren Rücken tragen werdet, und wahrlich ist es schlimm, womit ihr euch beladet.»"}
    KossonSILVER0576 = Kanon {"Kanon 20.  17", "Und sprich zu ihnen: «Ihr betrachtet das Leben in dieser Welt nur als Spiel und Zeitvertreib, und ihr wähnet, dass das Leben sei ein Würfelspiel."}
    KossonSILVER0577 = Kanon {"Kanon 20.  18", "Nicht ist dem aber so, denn es ist das Leben Weg und Ziel, fussend in der Evolution und im Werden zum SEIN (Schöpfungsleben), wollet ihr das denn nicht endlich begreifen?"}
    KossonSILVER0578 = Kanon {"Kanon 20.  19", "Ihr Menschen der Erde, ihr seid taub und stumm und blind in dichter Finsternis, und ihr wollet nicht hören die Wahrheit, und ihr wollet nicht reden über die Wahrheit, und ihr wollet nicht sehen die Wahrheit."}
    KossonSILVER0579 = Kanon {"Kanon 20.  20", "Ihr gehet selbst in die Irre, und ihr lasset euch auch treiben in die Irre von falschen Führern der Irre, die da sind Mächtige der Volksführung und die da sind Priester von Sekten und falschen Religionen mit Göttern und Götzen.»"}
    KossonSILVER0580 = Kanon {"Kanon 20.  21", "Und also rede weiter zu den Menschen in belehrender Form, sprichet der JHWH zum Propheten: «Was denket ihr, wenn eure selbsternannte Strafe für euer ungerechtes Leben über euch kommet?"}
    KossonSILVER0581 = Kanon {"Kanon 20.  22", "Und was denket ihr, wenn euch eure Stunde ereilet?"}
    KossonSILVER0582 = Kanon {"Kanon 20.  23", "Werdet ihr dann noch immer zu euren Göttern und Götzen von eigenen Gnaden rufen,oder die da fälschlich berufen sind durch euch selbst?"}
    KossonSILVER0583 = Kanon {"Kanon 20.  24", "Oder werdet ihr nach der Wahrheit rufen und nach der Schöpfung, die da euch wird Trost sein in der Stunde der Angst und Furcht?"}
    KossonSILVER0584 = Kanon {"Kanon 20.  25", "Wahrlich,so wird es sein,denn nach der Wahrheit und nach der Gerechtigkeit und nach der Schöpfung werdet ihr rufen in eurer Not."}
    KossonSILVER0585 = Kanon {"Kanon 20.  26", "Und wohl, wenn ihr ehrlich werdet und wenn euch euer Widerwille gegen die Wahrheit schmerzlich wird, falls ihr dazu noch imstande seid, dann möget ihr die Zeichen der Schöpfung und der Wahrheit und die Worte des JHWH und des Propheten wahrheitlich erkennen und also gerecht werden, so ihr nicht weiterhin der Ungerechten einer seid."}
    KossonSILVER0586 = Kanon {"Kanon 20.  27", "Es ist die Wahrheit geduldig im Ertragen, wie auch die Schöpfung ist geduldig, und so auch ist geduldig der JHWH und der Prophet, die da in unendlicher Geduld und Liebe belehrend und wartend sind, dass ihr findet und begehet den Weg des Gerechten.»"}
    KossonSILVER0587 = Kanon {"Kanon 20.  28", "Und belehre weiter die Unwissenden und die Ungerechten, und die Suchenden und Hoffenden, sprichet der JHWH zum Propheten: «Nur jene können die Wahrheit finden und aufnehmen und sie verstehen, die da den wahrheitlichen Worten des Propheten zuhören und die da die wahrlichen Zeichen sehen, die er ihnen aufzeiget und die da auch gegeben sind durch die Schöpfung oder durch den JHWH."}
    KossonSILVER0588 = Kanon {"Kanon 20.  29", "Alle jene aber sind sie tot im Leben, die da verpönen und verwerfen die Wahrheit und ihre Zeichen, denn sie leben schöpfungsfremd und sind abhold den Gesetzen und Geboten und der siebenheitlichen Ordnung der Schöpfung."}
    KossonSILVER0589 = Kanon {"Kanon 20.  30", "Und die lebenden Toten suchen andere Zeichen der Wahrheit und andere Worte der Wahrheit als die, die da sind wahrheitlich gegeben, denn sie suchen die Zeichen und die Wahrheit in Lügen und in Zauberei und Betrug, und so auch in magischem und mystischem Firlefanz."}
    KossonSILVER0590 = Kanon {"Kanon 20.  31", "Und es sind da jene die lebenden Toten, die da sagen und fordern, dass sie umgehend glaubend sein würden in der und jener Sache, wenn sie die Zeichen und die Wahrheit mit ihren Händen berühren könnten."}
    KossonSILVER0591 = Kanon {"Kanon 20.  32", "Aber es ist ihnen verwehret, die Zeichen zu berühren oder in ihre Nähe zu treten, weil ihr Sinnen und Trachten nach Zerstörung ist und nach Falschheit, so sie Gefahr verbreiten für alles Leben."}
    KossonSILVER0592 = Kanon {"Kanon 20.  33", "Und weil esihnen verwehret ist,die Zeichen zu berühren oder in ihre Nähe zu treten, so verfolgen und verleugnen sie die Zeichen und die Wahrheit."}
    KossonSILVER0593 = Kanon {"Kanon 20.  34", "Also ist ihnen gegeben nur der Weg,dass sie lernend sein müssen und erkennend in sich selbst And durch eigene Kraft, ohne dass sie die Zeichen berühren oder in ihre Nähe treten können, so sie allein wissend werden durch die Erforschung der Wahrheit in ihrem Innern."}
    KossonSILVER0594 = Kanon {"Kanon 20.  35", "Und es hat der JHWH die Macht, Zeichen hinabzusenden, und so auch die Seinen, die da sind seine Helfer, als Zeichen der Wahrheit; so auch die Schöpfung Zeichen hinabsendet in Alltäglichkeit und so zeuget von ihrer Existenz und Wahrheit."}
    KossonSILVER0595 = Kanon {"Kanon 20.  36", "Und es ist auch erarbeitet vom Propheten,Zeichen der Wahrheit zu tun,wie das getan ward von den Propheten ehedem und seit alters her."}
    KossonSILVER0596 = Kanon {"Kanon 20.  37", "Allein,es hat der Mensch der Erde sich gewandelt während der Zeit, die da ist dahingegangen; und er ist suchtvoll geworden nach Beweiskraft durch Sehen und Berührung und Hören."}
    KossonSILVER0597 = Kanon {"Kanon 20.  38", "Denn es ist dem Menschen entschwunden die Gesinnung der Bemühung, dass die Wahrheit soll erarbeitet werden durch eigene Kraft des Denkens und Forschens und Erkennens."}
    KossonSILVER0598 = Kanon {"Kanon 20.  39", "Also hat er sich gewandelt in der Form zum Irren, dass er nur noch das annehmen will als Wahrheit,das er mit seinen Händen berühren kann und das er mit seinen Augen sehen kann, und das er mit seinen Ohren hören kann."}
    KossonSILVER0599 = Kanon {"Kanon 20.  40", "Es ist dies aber der Weg der Falschheit und des reinen Materiellen, der jeder Bemühung des Selbstdenkens, des Selbstfühlens, des Selbstforschens, des Selbstsuchens und der Selbsterkennung jeglicher Lösung und Wahrheit Hohn spricht und spottet."}
    KossonSILVER0600 = Kanon {"Kanon 20.  41", "Also ist nicht gegeben dadurch der Weg zur Erforschung der Wahrheit und der Findung der Wahrheit, weshalb durch den JHWH und durch den Propheten keine Zeichen oder nur sehr sparsam Zeichen gegeben sein sollen, die da on der Wahrheit zeugend sind."}
    KossonSILVER0601 = Kanon {"Kanon 20.  42", "Gegeben ist das Wort der Wahrheit, und dieses ist genügsam vollauf, um die Lehre der Wahrheit zu offenbaren für jene, die da sind willig, die Mühsal der Erlernung durch eigene Kraft anzugehen und zu tragen.»"}
    KossonSILVER0602 = Kanon {"Kanon 20.  43", "Und es sprichet weiter der JHWH zum Propheten: <Belehre die Menschen, dass nicht sind alle Schätze des Wissens, des Könnens, der Liebe, der Wahrheit und der Weisheit der Schöpfung beim JHWH und beim Propheten, denn die Schätze und die Geheimnisse der Schöpfung sind unermesslich und niemals erfassbar für einen Menschen, denn die Schöpfung ist das unmessbare Geheimnis.»"}
    KossonSILVER0603 = Kanon {"Kanon 20.  44", "Und spreche, «dass nicht der JHWH und nicht der Prophet wissen um das Verborgene der Schöpfung, denn die Schöpfung allein weiss um ihr Verborgenes, und der Mensch vermag von ihr nur zu erfassen und zu erkennen das Unverborgene.»"}
    KossonSILVER0604 = Kanon {"Kanon 20.  45", "Und spreche belehrend zu den Menschen, «dass weder der JHWH noch der Prophet nahekommen der Schöpfung und dass weder der JHWH noch der Prophet sind Stellvertreter der Schöpfung."}
    KossonSILVER0605 = Kanon {"Kanon 20.  46", "Der JHWH und der Prophet sind Menschen von erfüllender Weisheit, Liebe und Wahrheit und von Können und Wissen, und über ihnen stehet unmessbar viel höher die Schöpfung als unmessbares Geheimnis."}
    KossonSILVER0606 = Kanon {"Kanon 20.  47", "Und es ist der JHWH ein Mensch der Erkenntnis in Weisheit, und so sei auch der Prophet, dem da offenbaret wird zuweilen vom JHWH die Belehrung in Wort und Schrift, die er da weitergebet und zu belehren hat an die Menschengeschlechter, für die er ist zugeordnet."}
    KossonSILVER0607 = Kanon {"Kanon 20.  48", "Und es ist weder der JHWH noch der Prophet strafend oder straffordernd, und sie üben weder Gewalt noch Zwang, noch belehren sie, dass geübet werden soll Gewalt oder Zwang."}
    KossonSILVER0608 = Kanon {"Kanon 20.  49", "Wahrlich,der JHWH entsendete Geistführer auf die Erde zu allen Menschengeschlechtern und Völkern, die da waren verfallen in Irrlehren zu Göttern und Götzen und die da lebeten in Irre und Wirrnis, so sie selbstverschuldet heimgesuchet waren durch Not und Drangsal,ehedem der Prophet der Neuzeit kam in eigener Pflichtberufung ein andermal im Wiederleben."}
    KossonSILVER0609 = Kanon {"Kanon 20.  50", "Und die Propheten waren zu ihnen gesendet, sie zu belehren und zu unterrichten in der Wahrheit, so sie frei werden sollten aus Drangsal und Not."}
    KossonSILVER0610 = Kanon {"Kanon 20.  51", "Doch die Völker der Erde bemüheten sich nicht um die Wahrheit, weder durch ihre Not und Drangsal, noch durch die Belehrung der Wahrheit durch die Propheten, die sie verleumdeten und verfolgeten."}
    KossonSILVER0611 = Kanon {"Kanon 20.  52", "Ihre Sinne waren verhärtet, so dem auch ist zur gegenwärtigen Zeit; und ihr böses und falsches Sinnen fern des wahrheitlichen Lebens liess sie alles Falsche weiterhin tun, was sie taten ehedem und das ihnen durch Irrlehren als wohlgetan erschien."}
    KossonSILVER0612 = Kanon {"Kanon 20.  53", "Und sie vergassen willentlich wieder die Lehre des JHWH und der Propheten, und sie vergassen alles willentlich, was ihnen an wahrheitlicher Lehre, Belehrungen und Ermahnungen war zuteil geworden."}
    KossonSILVER0613 = Kanon {"Kanon 20.  54", "Und sie öffneten sich mit Gewalt die Pforten zu allen Dingen, bis sie plötzlich erfasset wurden von den Folgen ihrer üblen Lebensweise, derweil sie sich noch erfreueten an allem, was sie sich übermütig in ungerechter Lebensweise selbst gegeben hatten."}
    KossonSILVER0614 = Kanon {"Kanon 20.  55", "Also fielen sie in Verzweiflung, und sie mordeten und töteten sich in Kriegen, und also erzeugeten sie Siechtum und Pestilenz, und sie brachten hervor falsche und tötende Gesetzgebungen und neuerliche Irrlehren also, die da waren dargebracht durch Priester und Machtgierige und durch Sekten und Religionen, die emporhoben in Eigenmacht Götter und Götzen."}
    KossonSILVER0615 = Kanon {"Kanon 20.  56", "Es waren aber verfolget und verleumdet die wahrheitlichen Propheten durch das Volk und durch die Ungerechten, und getreten wurden die Lehren und die Wahrheit."}
    KossonSILVER0616 = Kanon {"Kanon 20.  57", "Es waren aber gesendet durch den JHWH neue Geistführer, und also kamen in Eigenpflicht Propheten, die da stetig waren die Bringer und Künder der Wahrheit und wahrlich nur froher Botschaft."}
    KossonSILVER0617 = Kanon {"Kanon 20.  58", "Und es waren die Geistführer und Propheten auch Warner vor kommenden drohenden Geschehen und Warner vor falscher Lebensführung."}
    KossonSILVER0618 = Kanon {"Kanon 20.  59", "Und sie belehreten alle die Willigen und die Wahrheitssuchenden und so auch die Ungerechten, dass da keine Angst und keine Furcht und keine Trauer sein solle in ihnen, wenn sie die Wahrheit erkennen und zu leben beginnen nach der Wahrheit."}
    KossonSILVER0619 = Kanon {"Kanon 20.  60", "So aber ward von ihnen belehret, dass alle jene der selbstverschuldeten Strafe anheimfallen, die da ungehorsam sind, ungerecht und frevlerisch gegen die Gesetze und Gebote der Schöpfung und gegen die Wahrheit."}
    KossonSILVER0620 = Kanon {"Kanon 20.  61", "Und sei belehrend zu den Erdenmenschen, spricht der JHWH zum Propheten also; wollet ihr nicht endlich nachdenken und die Wahrheit erfassen, und wollet ihr euch nicht endlich der Wahrheit beugen?"}
    KossonSILVER0621 = Kanon {"Kanon 20.  62", "Treibet euch nicht selbst fort und immer weiter weg von der Wahrheit, die ihr sie wahrlich im tiefsten Innern eures Selbst suchet und erhoffet von morgens bis abends und die ihr sie verborgen im Innersten anrufet, im Trachten nach der wahrheitlichen Wahrheit der Schöpfung."}
    KossonSILVER0622 = Kanon {"Kanon 20.  63", "Und wahrlich, es suchet der Mensch die Wahrheit untergründig und in innerster Hoffnung, doch trotzdem lebet er in Ungerechtigkeit als Ungerechter, weil er irregeführt ist durch viele falsche Propheten von eigenen Gnaden,und von Priestern und Fanatikern der vielartigen Sekten und Religionen, die gesteuert sind von Gewaltherrschern und sonstigen Machtgierigen,von Gewinnsüchtigen, von Göttern und Götzen und Betrügern."}
    KossonSILVER0623 = Kanon {"Kanon 20.  64", "Wahrlich,Frieden und Liebe sei mit den Gerechten, und Frieden und Liebe sei mit jenen, die da die Wahrheit der Schöpfung suchen."}
    KossonSILVER0624 = Kanon {"Kanon 20.  65", "Die Schöpfung hat es bestimmet, dass Frieden, Liebe und Barmherzigkeit erwiesen werden jedem,der da unwissentlich etwas Böses tuet oder sonstig etwas wider die Gesetze und Gebote oder gegen die siebenheitliche Ordnung, wenn daraufhin folget Erkenntnis und Reue und der Weg der Wahrheit beschritten wird."}
    KossonSILVER0625 = Kanon {"Kanon 20.  66", "Die Schöpfung ist allvergebend und barmherzig, und also machet sie auch klar ihre Existenz und Liebe so,dass da der Fehlbare sie und die Wahrheit erkenne und nach ihren Gesetzen und Geboten lebe."}
    KossonSILVER0626 = Kanon {"Kanon 20.  67", "Und die Zeichen der Liebe der Schöpfung beweisen, dass bei ihr ist Vergebung für alles Fehlbare, wenn das Fehlbare wird behoben und das Leben des Rechtens wird."}
    KossonSILVER0627 = Kanon {"Kanon 20.  68", "Und die Zeichen der Liebe der Schöpfung beweisen, dass da ist irrig und falsch alles im Tun, wenn angebetet (geehrt) werden andere Dinge nebst der Schöpfung, wie sie da sind Götter und sonstige Menschen und Götzen also."}
    KossonSILVER0628 = Kanon {"Kanon 20.  69", "Also soll der Mensch den bösen Gelüsten und Irrlehren nicht folgend sein, so er nichts anderes neben der Schöpfung verehret und anbetet (ehrt)."}
    KossonSILVER0629 = Kanon {"Kanon 20.  70", "Achtet der Mensch aber nicht dieser Lehre und betet (bettelt und fleht) er an oder verehret er Götter oder Götzen oder Menschen, dann gehet er wahrlich irre und ist nicht der Rechtgeleiteten einer."}
    KossonSILVER0630 = Kanon {"Kanon 20.  71", "Den Gerechten und Wissenden ist es untersaget durch Vernunft und Verstand, anzurufen und anzubeten (anzuflehen und zu ehren) Götter und Götzen, so sie nicht folgend sein können den bösen Gelüsten der Verführung zum Falschen und Irren."}
    KossonSILVER0631 = Kanon {"Kanon 20.  72", "Der JHWH und der Prophet berufen sich auf den klaren Beweis der Schöpfung, doch ihr, ihr verwerfet den Beweis, weil ihr euch wissender wähnet und weiser als dies da ist die Schöpfung selbst, oder als dies jemals tun und sein können für sich selbst der JHWH und der Prophet und die Geistführer."}
    KossonSILVER0632 = Kanon {"Kanon 20.  73", "Als Beweis wünschet ihr Zeichen und Wunder, doch gerade die sollen euch verschlossen und abgewendet bleiben, denn verstehen würdet ihr sie einerseits nicht, und verleugnen würdet ihr sie andererseits in eurem Unverstand und in eurer Ungerechtigkeit."}
    KossonSILVER0633 = Kanon {"Kanon 20.  74", "So seien spärlich gegeben in Ausnahmen die Zeichen der Wahrheit durch den JHWH und durch den Propheten nur jenen, die da wahrlich Gerechte sind oder die da sind werdende Gerechte."}
    KossonSILVER0634 = Kanon {"Kanon 20.  75", "Wahrlich, nur den Gerechten und den werdenden Gerechten sind die Zeichen der Wahrheit Beweisführung für die Wahrheit, so auch sie die Zeichen der Beweisführung sehen sollen in spärlicher Form, die da gegeben sind von Zeit zu Zeit durch den JHWH und den Propheten."}
    KossonSILVER0635 = Kanon {"Kanon 20.  76", "Bei der Schöpfung sind die Schlüssel des Verborgenen, und das Verborgene kennt nur die Schöpfung allein."}
    KossonSILVER0636 = Kanon {"Kanon 20.  77", "Und was das Verborgene der Schöpfung ist, das ist nicht erfassbar und erkennbar durch den Menschen, und so also nicht durch den JHWH und nicht durch den Propheten."}
    KossonSILVER0637 = Kanon {"Kanon 20.  78", "Also können der JHWH und der Prophet nur erforschend, erkennend und belehrend sein in jenen Dingen, die da sind durch Menschen erkennbar und erklärbar, und die da können erkläret werden dem JHWH und dem Propheten durch Wesenheiten des Lichtes aus ihnen zugänglichen Sphären des Geistes, und die da sind unzugänglich allen Menschen, die da nicht lebend sind im Wissen, in der Liebe, in der Weisheit und in der Wahrheit jener Ebenen, in denen da lebend sind der JHWH und der Prophet."}
    KossonSILVER0638 = Kanon {"Kanon 20.  79", "Und was das Verborgene ist bei der Schöpfung, das ist ihr unmessbares Geheimnis, das da erst wird erkennbar und erfassbar und verstehend im Geiste des Menschen, wenn er dereinst eingehet in das SEIN (Schöpfungsleben) der Schöpfung."}
    KossonSILVER0639 = Kanon {"Kanon 20.  80", "Und es ist wahrlich die Schöpfung, in die der Mensch dereinsten findet die Heimkehr, in früherer oder späterer Zeit, gemessen je nachdem, wie da waren sein Leben und seine Werke."}
    KossonSILVER0640 = Kanon {"Kanon 20.  81", "Denn es ist gegeben, dass der Mensch sein Leben friste und führe in Gerechtigkeit und in weiser Erfüllung der Gesetze und Gebote der Schöpfung, so er wandle in stetiger Evolution, um zu gestalten seine Bemühungen der gerechten Lebensführung zur relativen Vollkommenheit, so er sich zur Geistform wandle und zur späteren relativen Vollkommenheit, die da die Heimkehr gewähret in die Schöpfung als Endziel."}
    KossonSILVER0641 = Kanon {"Kanon 20.  82", "Und wahrlich, es ist die Schöpfung das Leben und das Höchste über allem Leben, und also stehet sie über dem Leben und über dem Tod, wie sie auch stehet als Endziel aller Evolution, so als wahrliches SEIN (Schöpfungsleben), zu dem der Mensch durch seine Evolution in unaufhaltsamer Bemühung hinstrebe."}
    KossonSILVER0642 = Kanon {"Kanon 20.  83", "Und wahrlich, gegeben ist durch die Schöpfung alles Leben und alle höchste erreichbare Evolution, so also die Liebe, die Wahrheit, das Wissen, die Weisheit und die Folgerichtigkeit und auch alles Erforderliche an Nahrung und Umgebung im Lebensteil des materiellen Bereiches."}
    KossonSILVER0643 = Kanon {"Kanon 20.  84", "Also ist da gegeben durch die Schöpfung gesamtheitlich alles das, was da der Mensch bedarf zum Leben und zur Entwicklung, doch dennoch weichet er ab von der Schöpfung und ihrer Wahrheit und stellet ihr zur Seite Götter und Götzen, die da über ihn nur bringen Not, Trübsal und Drangsal und irrige Lehren in unendlicher Zahl."}
    KossonSILVER0644 = Kanon {"Kanon 20.  85", "Und es bringet der Menschen Glaube an Götter und Götzen das Übel in grossen Massen über die Völker und Menschengeschlechter, denn es werden durch den Glauben und durch Götter und Götzen, durch Priester, Machtgierige, Herrschende und Gewinnsüchtige falsche Gesetze und Gebote ersonnen, die da widersinnig sind zu den Gesetzen und Geboten der Schöpfung."}
    KossonSILVER0645 = Kanon {"Kanon 20.  86", "Und es bringen die falschen Gesetze und Gebote der Sekten und Religionen und Machtgierigen und Gewinnsüchtigen den Völkern und Menschengeschlechtern nur Gewalt und Zwang und Knechtschaft und Schaden, denn sie sind ersonnen allesamt nicht zur wahrheitlichen Führung des Menschen, sondern zu seiner Versklavung und zur zwingenden Gewaltstrafe bei schon nur geringstem Anlass."}
    KossonSILVER0646 = Kanon {"Kanon 20.  87", "Also bringen die falschen Gesetze und Gebote der Sekten und Religionen und der in Gewaltsamkeit Regierenden und Machtgierigen und Gewinnsüchtigen nicht Gerechtigkeit und Rechtschaffenheit und Freiheit, sondern nur Strafgericht, Gewalttat, Krieg, Siechtum und Tod und mannigfache Übel, so auch der Mensch im einzelnen wird gepräget dadurch."}
    KossonSILVER0647 = Kanon {"Kanon 20.  88", "Und die Machtgierigen, die Gewinnsüchtigen, die Regierenden und die Sekten und Religionen und ihre Priester spalten das Einheitssinnen der Menschen in Uneinigkeit in allen Dingen, und sie stiften Verwirrung und Feindschaft, so der eine wider den andern ist, und so der Nächste seinem Nächsten Gewalt antut, und so der Sohn wider den Vater und die Tochter wider die Mutter in Feindschaft verfallen."}
    KossonSILVER0648 = Kanon {"Kanon 20.  89", "Und die Ungerechten, die Machtgierigen, die Gewinnsüchtigen, die Herrschenden und die Regierenden, und die Priester und Sekten und Religionen treiben Kraft ihrer Gewalt und ihrer falschen Gesetze und Gebote die Menschen zusammen zu Kriegsheeren, um tödliche Schlachten und Kriege zu führen in Machtgier und Hass, wider die Brüder und Schwestern und wider die Kinder des Nächstenvolkes."}
    KossonSILVER0649 = Kanon {"Kanon 20.  90", "Doch es achtet der Mensch nicht dieser Wahrheit, und er verwirfet die Wahrheit weiterhin, und wenn er die Zeichen der Wahrheit siehet und die Worte der Wahrheit höret, dann sprichet er nur töricht und verleumdend über sie."}
    KossonSILVER0650 = Kanon {"Kanon 20.  91", "Es ist aber geweissaget von alters her, dass die Verwerfung der Wahrheit den irdischen Menschengeschlechtern und Völkern bringen wird Übel und Leid und Trauer und dass die Menschen fallen werden in Not und Drangsal und Trübsal, und wahrlich ward dem so in stetiger Folge, wie dem ward geweissaget."}
    KossonSILVER0651 = Kanon {"Kanon 20.  92", "Es ist gesetzet für jede Weissagung eine bestimmte Zeit, so sie sich dann erfüllen wird, und so ist dies auch gegeben für die Weissagungen der Neuzeit, die da sind niedergeschrieben seit alters her und die auch sind neuerlich gegeben durch den Propheten der Neuzeit, und bald werdet ihr es erfahren und die Schrecknisse und die Tode erdulden müssen, denn es sind gesetzet und erfolget in Erfüllung die Zeichen der Zeit, die da künden von der Endzeit der Erfüllung der Weissagungen."}
    KossonSILVER0652 = Kanon {"Kanon 20.  93", "Doch es lebet der Mensch der Erde weiterhin in falschem Glauben, Sinnen und Trachten und Handeln, wie alles ist fälschlich gegeben durch die irren Gesetze und Gebote von Sekten und Religionen, und von Herrschsüchtigen und falschen Regierenden, so der Erdenmensch allem Wissen und aller Wahrheit abtrünnig und fremd ist wie seit alters her."}
    KossonSILVER0653 = Kanon {"Kanon 20.  94", "Es ward aber belehret die Wahrheit des wahrlichen Wissens den irdischen Menschengeschlechtern seit alters her und schon gegeben dem Adam."}
    KossonSILVER0654 = Kanon {"Kanon 20.  95", "Es waren gegeben den Menschengeschlechtern und Völkern der Erde Propheten von alters her, so also ward gesendet der Henoch und der Elja, der Jesaja und der Jeremja, und der Jmmanuel und der Mohammed in direkter Folge und steter neupersönlicher Wiedergeburt, nebst dem Johannes und Eljas und dem Hjob und allem Heer der anderen Rechtschaffenen und Gerechten, wie sie da auch waren der Buddha, der Zoroaster und der Babatschi und andere."}
    KossonSILVER0655 = Kanon {"Kanon 20.  96", "Sie waren alle erkoren und geleitet und hinabgesendet, teilens durch den JHWH und teilens durch sich selbst und im Wissen und Können als Propheten oder Geistführer, zu den Menschengeschlechtern und Völkern der Erde, doch alle wurden sie verleumdet und verfolget, und es wurden ihre Lehre und ihre Worte der Wahrheit im Unverstehen oder Hass verfälschet, so entstanden sind daraus irreführende und falsche Religionen und Sekten."}
    KossonSILVER0656 = Kanon {"Kanon 20.  97", "Und ein andermal ist gesendet nun ein Prophet hinab zur Erde in eigener Pflichterfüllung, und ein andermal ist gegeben das Buch des Wortes der Wahrheit, das da offenbaret wird in verständiger Form, voll des Segens und der wahrheitlichen Richtlinien und um zu erfüllen das Frühere, und auf dass belehret werden die irdischen Menschengeschlechter und Völker, und auf dass endgültig verwarnet seien die Fehlbaren und die Ungerechten."}
    KossonSILVER0657 = Kanon {"Kanon 20.  98", "Und es sei letztliches Mal und für alle Zeit der Zeit, dass gesendet werde ein Prophet von dieser Grösse und der ist gesetzgebend, und dass gegeben werde das Buch des Wortes der Wahrheit und die Lehre des Geistes also an die irdischen Menschengeschlechter, und niemals soll Gleiches wiederkehren in alle Zukunft."}
    KossonSILVER0658 = Kanon {"Kanon 20.  99", "Gegeben ist die Wahrheit und der Prophet, und dies sei von Gültigkeit nunan für alle Zeit des Kommenden, denn zerstöret wurden alle Bestimmungen durch die Menschen der Erde, die da hätten anderes ermöglicht und ein also weiteres Wirken des Propheten in Zukunft."}
    KossonSILVER0659 = Kanon {"Kanon 20. 100", "Also ist nunmehr gegeben der letzte Prophet diese Wichtigkeit für alle Zeit, und so auch das Buch des Wortes der Wahrheit ein letzliches Mal, und nicht wird folgen eine weitere Schrift und ein Prophet gleicher Wichtigkeit ein andermal, so aber wird auch nicht folgen ein Reformer, denn gegeben ist alles im Buch des Wortes der Wahrheit und durch den Propheten und dessen Lehre des Geistes in erschöpfender Form, so Reformer und eine neue Lehre und ein neuer Prophet nach dem Propheten der Neuzeit nie und niemals sein werden, und so doch welche kommen und sich Gesandte oder Reformer nennen in eigenen Gnaden, dann die also voller Falschheit sind und voller Irrung, Lug und Trug."}
    KossonSILVER0660 = Kanon {"Kanon 20. 101", "Also wird nicht sein ein andermal ein Prophet nach dem Propheten der Neuzeit, und so wird auch nicht sein ein Reformer oder Gesandter anderer Art, denn gegeben ist die Wahrheit der wahrlichen Wahrheit durch den Propheten der Neuzeit, und es ist die gegebene Wahrheit aller Wahrheit die vollumfängliche Wahrheit, die nicht kann reformiert werden in auch nur einem Tüpfelchen."}
    KossonSILVER0661 = Kanon {"Kanon 20. 102", "Und dies ist wahrheitlich das Buch des Wortes der Wahrheit, das da ist offenbaret durch den JHWH und den Propheten, und das da ist unveränderbar und nicht reformbedürftig für alle Zeit."}
    KossonSILVER0662 = Kanon {"Kanon 20. 103", "Es ist das Buch des Wortes der Wahrheit voll des Segens, zu erfüllen das frühere und wahrliche Leben, auf dass da alle Menschen rund auf der Erde verwarnet und belehret seien."}
    KossonSILVER0663 = Kanon {"Kanon 20. 104", "Und es ist dies das Buch des Wortes der Wahrheit für alle, so für die Gerechten und für die Ungerechten also, und es ist in keinem Tüpfelchen veränderbar oder zu reformieren, weder durch Machtgierige und Regierende noch durch Priester von Sekten und Religionen, und nicht durch Gewinnsüchtige und Zweifler, und also nicht durch falsche Gesandte und falsche Reformer und falsche Propheten."}
    KossonSILVER0664 = Kanon {"Kanon 20. 105", "Und es ist gegeben das Buch des Wortes der Wahrheit zur Belehrung der Ungerechten und zur Wegeleitung jener, die da bereits im Ganzen oder in Teilen um die Wahrheit wissend sind und sich streng halten im Gesetz und Gebot zur Schöpfung."}
    KossonSILVER0665 = Kanon {"Kanon 20. 106", "Wer aber ist ungerechter, als wer Lügen ersinnet wider die Schöpfung, wie da getan wird durch die Priester von Sekten und Religionen, und von falschen Reformern, und von falschen Gesandten und falschen Propheten und Machtgierigen und Gewinnsüchtigen, die da lügen und sprechen, dass ihnen offenbaret worden wäre dieses und jenes, während ihnen doch nichts offenbaret worden ist."}
    KossonSILVER0666 = Kanon {"Kanon 20. 107", "Und wer ist ungerechter, als wer Lügen ersinnet wider die Wahrheit und wider die Gesetze und Gebote und die siebenheitliche Ordnung der Schöpfung, wie da getan wird durch Priester von Sekten und Religionen, durch Betrüger und Ausgeartete und durch Gewinnsüchtige und Machtgierige, die da lügen und sprechen, dass ihnen offenbaret worden wäre dieses und jenes von Verstorbenen aus dem Totenreiche oder von Geistformen und Dämonen von höheren oder niederen Ebenen, während ihnen doch nichts offenbaret worden ist, weil dies für sie nicht etwas des Möglichen wäre, so sie daran lügen und betrügen."}
    KossonSILVER0667 = Kanon {"Kanon 20. 108", "Sie aber lügen und betrügen, um sich bei den Mitmenschen und bei ihren Gläubigen ihrer Irrlehren in Grösse erscheinen zu lassen, und um auch zu befriedigen ihre Gelüste der Selbstsucht und der Selbstemporhebung im Bedürfnis der Geltung, weil sie da wahrlich sind Ungerechte und Unfähige zur rechten Lebensführung, was sie aber verbergen wollen durch ihre Lügen und durch Betrug."}
    KossonSILVER0668 = Kanon {"Kanon 20. 109", "Und sie lügen und betrügen alle im Unwissen der Wahrheit, dass da die Gesetze und Gebote der Schöpfung nicht und niemals solches zulassen und es also verunmöglichen, dass Menschen können in Verbindung treten mit Verstorbenen oder mit Geistformen oder Dämonen von höheren oder niederen Ebenen, wenn nicht ist erarbeitet und erreichet durch den Menschen der Stand der universellen Liebe und der universellen Wahrheit, und so also das universelle Wissen und die universelle Weisheit, wie diese gegeben sein müssen dem JHWH und den Propheten."}
    KossonSILVER0669 = Kanon {"Kanon 20. 110", "So der Mensch aber erlanget und erreichet hat dieses Können, dann ist er wahrheitlich wissend geworden, dass die Toten nicht wissender sind als die Lebenden, so diese nicht geben können den Lebenden Rat oder Belehrung im Wissen oder über Unbekanntes, und so der wahrlich Könnende dies weiss, so enthaltet er sich jeder Verbindung und Verständigung mit Verstorbenen, und er offenbaret auch nicht das Geheimnis seines Könnens den Menschen."}
    KossonSILVER0670 = Kanon {"Kanon 20. 111", "Und so der Mensch erlanget hat dieses Können, so erschliessen sich ihm die nächsten sieben höheren Ebenen der Geistformen und des Geisteslebens, die ausserhalb seines materiellen Lebens für ihn belehrend sind, doch er verwahret sich dagegen, eine der höheren Ebenen selbst anzurufen, weil die Kommunikation nur erfolget von der höheren zur niedrigeren Ebene in belehrender und seltener Form, und nur gegeben und bestimmet durch die höhere Ebene."}
    KossonSILVER0671 = Kanon {"Kanon 20. 112", "Und es sind nicht und niemals unter dem Volke Menschen dieses Könnens, die da vermöchten zu tun solchermassen, und die, die da sind die wenigen in Befähigung unter den Menschengeschlechtern, die leben in Schweigsamkeit und in Einsamkeit, so sie nicht bedränget und befraget oder verleumdet und verfolget werden durch die Unwissenden und durch die Lügner und Betrüger."}
    KossonSILVER0672 = Kanon {"Kanon 20. 113", "Also ist gegeben die Wahrheit, dass da alles sind Lügner und Betrüger, die sich offentun als Beschwörer und Kommunizierer mit Verstorbenen und mit Geistern und Dämonen, und sie alle sind da Lügner und Betrüger, die sich da geben in dieser Form als Magier und Medium, oder als Weissager, wie auch alles sind Lügner und Betrüger, die Magisches und Mystisches tun in Kulten und mit allerhand Tand und Hilfsgerätschaften und sich in dieser Form nennen Magier, Medium und Weissager oder Heiler."}
    KossonSILVER0673 = Kanon {"Kanon 20. 114", "Es sind da aber welche unter den Menschen der Völker, die da haben wahrliche Gesichte im Schlaf und im Wachen, und die da sind genennet als Seher und Hellsichtiqe und die da wahrheitlich schauen Geschehen der Vergangenheit und Zukunft in Hellsichtigkeit."}
    KossonSILVER0674 = Kanon {"Kanon 20. 115", "Die da aber haben Gesichte, deren sind nur wenige unter jedem Volke in wahrlich wahrheitlicher Form, wenn sie sind frei vom Glaubenssinnen und frei von Sekten und Religionen, wenn sie sind wahrheitlich Wissende um die Wahrheit der Schöpfung und um ihre Gesetze und Gebote und ihre siebenheitliche Ordnung."}
    KossonSILVER0675 = Kanon {"Kanon 20. 116", "So sie dem aber nicht sind, so fügen sie den Werten der Gesichte eigene Unwerte hinzu und werden dadurch Verratende an der Wahrheit, wie die Lügner und Betrüger Verratende sind an der Wahrheit, die da Gesichte erlügen als Offenbarung, die da ihnen aber nicht offenbaret sind, wie das ist gegeben in grosser Zahl von Lügnern und Betrügern in Selbstsucht, Gewinnsucht und Geltungssucht."}
    KossonSILVER0676 = Kanon {"Kanon 20. 117", "Und viele sind jener, die da fälschlich leben und lügen und betrügen und die da sagen und sprechen, dass sie könnten herabsenden Zeichen gleich der Schöpfung und gleich dem JHWH und gleich dem Propheten."}
    KossonSILVER0677 = Kanon {"Kanon 20. 118", "Doch sie können es tun mitnichten und nur durch Lügen und Betrug, so durch Taschenspielerei, die da ist genannt Zauberei und Illusion, oder aber durch andere Täuschung irgendwelcher Art, so also auch durch gedankliche Macht, durch die da vorgegaukelt wird dieses und jenes, das da weder ist gegenwärtig noch existent."}
    KossonSILVER0678 = Kanon {"Kanon 20. 119", "Also sind sie Frevler an der Wahrheit, und also sind sie Verlierer und Verlorene, die da heulen und haben Zähneklappern im Leben bei jeder Gelegenheit, so aber auch im Sterben und im Tod und im Wiederleben, und also birget der Tod für sie Schrecken und Entsetzen, und es wird ihr eigenes Strafgericht für sie zu Schlünden des schrecklichen Grauens."}
    KossonSILVER0679 = Kanon {"Kanon 20. 120", "Sie bringen Schande über sich selbst, und Schande und Grauen ist der Lohn ihres Tuns, was sie da alles Falsches getan und gesprochen haben wider die Wahrheit und die Schöpfung, und weil sie sich abgewendet haben in Verachtung von der wahrheitlichen Wahrheit und von den Gesetzen und Geboten der Schöpfung."}
    KossonSILVER0680 = Kanon {"Kanon 20. 121", "Und all die Ungerechten,die Frevler und die Lügner und Betrüger, sie verfallen ihrem eigenen Strafgericht, nackend und bloss, wie sie waren erschaffen, und alles hinter sich lassend, was ihnen durch die Schöpfung ward gegeben; und alles hinter sich lassend, was sie sich da haben gehortet durch Ungerechtigkeit und Falschheit und durch Frevel, Lügen und Betrug."}
    KossonSILVER0681 = Kanon {"Kanon 20. 122", "Sie werden allein sein in ihrer Stunde des Grauens und hinabsinken in die brüllenden und lodernden Schlünde ihres Entsetzens, und nicht wird sein für sie ein Fürsprecher, den sie da wähnten unter ihren Göttern und Götzen, denn niemals bringeten Lüge und Betrug und Irrlehre und Ungerechtigkeit Fürsprache in Not und Drangsal."}
    KossonSILVER0682 = Kanon {"Kanon 20. 123", "Wahrlich, allein die Gesetze und Gebote der Schöpfung und ihre siebenheitliche Ordnung lassen keimen das Korn und den Samen der selbst winzigsten Pflanze."}
    KossonSILVER0683 = Kanon {"Kanon 20. 124", "Die Gesetze und Gebote der Schöpfung und ihre siebenheitliche Ordnung allein bringen das Lebendige hervor aus dem Ruhenden, das da dem Menschen in Unwissenheit als tot erscheinet, das da aber wahrheitlich Übergang und Wandel zum Wiederleben ist."}
    KossonSILVER0684 = Kanon {"Kanon 20. 125", "Und die Gesetze und Gebote der Schöpfung und ihre siebenheitliche Ordnung allein bringen da Lebendiges in das Ruhende, das da dem unwissenden Menschen als tot erscheinet, das da aber ist wahrheitlich Übergang und Wandel zum Wiederleben."}
    KossonSILVER0685 = Kanon {"Kanon 20. 126", "Und es könnten Wasser sein und Wind und Schnee, und es könnten sein Regen, Hagel und Sturm und Jahreszeiten, und es könnten sein Sonnenschein, Licht, Dunkelheit und Wärme und Kälte, und doch würden nicht sein Werden, Leben und Vergehen, und es wären nicht Geburt und nicht Tod, kein Keimen, kein Gedeihen und kein Blühen, und es wären keine Frucht und keine Ernte, wenn da nicht wären gegeben und bestimmend die Gesetze und Gebote der Schöpfung und ihre siebenheitliche Ordnung, damit da alles sei werdend, lebend und vergehend, was da alles ist existent an allen Heeren des Lebens, die da sind erschaffen durch die Schöpfung, der Weisen, der Gerechten."}
    KossonSILVER0686 = Kanon {"Kanon 20. 127", "Es liess die Schöpfung durch ihre Gesetze und Gebote den Tag anbrechen zum Leben, und die Nacht zur Ruhe, und die Sonne und die Gestirne für die Wärme und für das Licht und die Kälte und für die Bewegung, und also die Zeit für das Werden und Leben und Vergehen und für die Evolution, und den Raum und die Feste und die Wasser und Lüfte für die Existenz aller Heere des Lebens."}
    KossonSILVER0687 = Kanon {"Kanon 20. 128", "Und es ist dies die Ordnung der Schöpfung durch ihre Gesetze und Gebote und durch die Siebenheit, und Preis (Ehre) und Dank sei der Schöpfung, die da ist wahrliche Liebe und wahrliche Weisheit, und die da ist wahrliches Wissen und wahrliche Wahrheit in endloser Geduldigkeit und Langmut und Allmächtigkeit."}
    KossonSILVER0688 = Kanon {"Kanon 20. 129", "Und es sind bis in die letzte Einzelheit die Zeichen der Schöpfung dargeleget allen Heeren des Lebens und dem Menschen, auf dass sie erkennet und befolget und als Hilfe genommen werden zur gerechten Lebensführung, wie der Mensch da auch nimmt die Sterne am Himmel als Hilfe, zur Einhaltung der rechten Richtung in der Finsternis des Landes und auf den wogenden Wassern der Meere und in den Lüften über der Erde und in der luftlosen Leere des Weltenraumes."}
    KossonSILVER0689 = Kanon {"Kanon 20. 130", "Und es sind bis ins einzelne dargeleget und gegeben die Zeichen der Schöpfung für alle Heere des Lebens und für den Menschen, damit der Mensch die Wahrheit erkenne und befolge und den schöpferischen Gesetzen und Geboten folgsam sei, wissend und weise und beachtend und ehrend."}
    KossonSILVER0690 = Kanon {"Kanon 20. 131", "Und es ist die Schöpfung, die da hat durch ihre Gesetze und Gebote und durch ihre siebenheitliche Ordnung euch kreieret aus winzigsten Elementen der Erde und aus winzigsten Verbundenheiten gewandelten Lebens aus Fauna und Flora."}
    KossonSILVER0691 = Kanon {"Kanon 20. 132", "Und es ist euch gegeben der Ort eurer Kreation zum Verweilen und zur Evolution in Befolgung der schöpferischen Gesetze und Gebote und der siebenheitlichen Ordnung, so also eure Erde eure Welt ist und auf der ihr erfüllen sollet alle Richtlinien des Gerechten in Folgerichtigkeit."}
    KossonSILVER0692 = Kanon {"Kanon 20. 133", "Und also sind bis ins einzelne dargeleget alle Zeichen der Schöpfung für alle Menschen, und sie sind gegeben in allem und jedem und für alle Heere des Lebens also, auf dass alles Leben und der Mensch sie erkenne und begreife und befolge, und auch sollet ihr sie endlich erkennen, begreifen und befolgen."}
    KossonSILVER0693 = Kanon {"Kanon 20. 134", "Und es sind da die Gesetze und Gebote der Schöpfung und ihre siebenheitliche Ordnung,die da den Regen als Wasser aus den Wolken niederfallen lassen, so dadurch wird Wachstum hervorgebracht und also Grünendes, woraus gereihtes Korn spriesset, so aber auch Pflanzen jeglicher Art und Gesträuch und Bäume aller Art, und die da vollbehangen sind mit wohlig duftenden Blütendolden, aus denen Früchte erwachsen in unendlicher Zahl und in mannigfacher Form und in ähnlicher oder unähnlicher Art."}
    KossonSILVER0694 = Kanon {"Kanon 20. 135", "Und wahrlich, betrachtet ihr die Blüten und Früchte der Pflanzen, Gesträuche und Bäume, wenn sie ihre Blüten und Früchte hervorbringen, und wie sie reifen und neue Samen bringen, dann erkennet ihr darin wahrlich die Zeichen der unendlichen Liebe der Schöpfung, und das Wissen und die Wahrheit und die Weisheit und die Geduld, die da allesamt sind erkenntlich und verankert in den Gesetzen und Geboten und der siebenheitlichen Ordnung der Schöpfung."}
    KossonSILVER0695 = Kanon {"Kanon 20. 136", "Es sind aber sichtbar und erkenntlich alle Zeichen der Schöpfung nur jenen, die da wahrlich sind Wissende und Gerechte oder solche, die da sind in Ehrlichkeit gehend auf dem Wege zum Wissen und zur Gerechtigkeit und zur Wahrheit und denen da ist fern jeglicher Glaube, der da nur ist unwissende Annahme, weil nur wissende Wahrheit wahrliche Wahrheit und wahrliches Wissen in sich birget."}
    KossonSILVER0696 = Kanon {"Kanon 20. 137", "Doch es sind nur wenige Wissende und so also wenige Gerechte, und viel mehr sind der Unwissenden und der Ungerechten, die da sich selbst oder Nebenbuhler und Götter und Götzen zur Seite der Schöpfung stellen, obwohl sie da ist die Alleinige, die da in ihrem Himmel (Universum) hat alles Leben erschaffen und alle Existenz."}
    KossonSILVER0697 = Kanon {"Kanon 20. 138", "Und die Ungerechten und Unwissenden dichten der Schöpfung Nebenbuhler bei, die da sind Götzen und Götter und dergleichen und die da sind auch ein einzelner Gott oder ein einzelner Götze."}
    KossonSILVER0698 = Kanon {"Kanon 20. 139", "Und also dichten die Ungerechten und die Unwissenden fälschlich der Schöpfung Söhne und Töchter bei und Engel und Heilige (Ehrwürdige), so sie dies tun können bei ihrem Gott oder Götzen, oder bei ihren vielen Göttern oder Götzen und dergleichen."}
    KossonSILVER0699 = Kanon {"Kanon 20. 140", "Es ist die Schöpfung aber heilig (ehrwürdig) und erhaben über alle Lügen und Andichtungen, und es sind so auch heilig (ehrwürdig) und erhaben alle ihre Gesetze und Gebote und ihre siebenheitliche Ordnung."}
    KossonSILVER0700 = Kanon {"Kanon 20. 141", "Und also ist die Schöpfung die Schöpfung und die höchste Macht, und sie stehet in unmessbarer Grösse und Ganzheit über allen ihren Kreationen, die da sind Schöpfungen der Schöpfung und die da sind Mensch und Tier und Pflanzen aller Art, und Welten und Sonnen, Galaxien, Raum und Zeit, Licht und Finsternis und was da alles ist an sonstigen Heeren des Lebens jeglicher Art."}
    KossonSILVER0701 = Kanon {"Kanon 20. 142", "Und die Schöpfung ist ohne Geschlecht, und sie ist da weder Mann noch Weib, noch keusch oder unkeusch, und weder zeitlich noch unzeitlich, oder weder räumlich noch unräumlich, denn die Schöpfung ist wahrheitlich Geist und Geistkraft und Licht und wahrlich reine Geistenergie in bewusster Form."}
    KossonSILVER0702 = Kanon {"Kanon 20. 143", "Und es ist die Schöpfung ohne Gefährten und ohne Gefährtin, denn die Schöpfung ist ES und geschlechtslos in wahrheitlicher Form, so sie sich also unterscheidet in allem und von jeglichem Leben."}
    KossonSILVER0703 = Kanon {"Kanon 20. 144", "Und so die Schöpfung ist ohne Geschlecht und ohne Gefährten oder Gefährtin, und da sie ist Geistenergie in reinster wahrheitlicher Form und ohne Bindung in abhängiger Weise, wie sollte sie da haben Söhne und Töchter oder solche sich berufen in irgendwelcher Art und Form, und wie könnte sie da sein ein Gott, der doch nur ist ein Mensch und Gewaltherrscher von grobmaterieller Lebensart."}
    KossonSILVER0704 = Kanon {"Kanon 20. 145", "Die Schöpfung ist die Schöpfung aller Dinge, und es gibt keine Schöpfung in ihrem Himmel （Universum) ausser ihr, und sie allein soll verehret und angebetet (geehrt) werden, die da allein ist wahrheitlich die Schöpfung und die Hüterin aller Heere des Lebens und alles Existenten, das da zudem ist wahrheitlich erschaffen von der Schöpfung allein."}
    KossonSILVER0705 = Kanon {"Kanon 20. 146", "Wahrlich, sie da ist die Schöpfung, die da ist eure Schöpfung, und deren Blicke und SEIN （Schöpfungsleben) ihr als Menschen niemals erreichen und sehen und erfassen könnet, die sie aber erreichet und sieht und erfasset euren Blick und euer Leben, wie immer ihr da auch blicket und wie ihr auch führet euer Leben zu aller Zeit."}
    KossonSILVER0706 = Kanon {"Kanon 20. 147", "Und die Schöpfung vermag alles, denn sie ist das unmessbare Geheimnis, und sie ist das Unbegreifliche,das Allheitliche und das Allkundige, und sie ist dem Menschen als eigene Existenz das Verborgenste alles Verborgenen, so also das unmessbare Geheimnis."}
    KossonSILVER0707 = Kanon {"Kanon 20. 148", "Wahrlich, das ist die Schöpfung, die da ist eure Schöpfung, durch die da offenbaret ist durch Zeichen in mannigfacher Form die siebenheitliche Ordnung und die Gesetze und Gebote des Lebens, der Evolution und des SEINs (Schöpfungslebens) für alle Heere des Lebens."}
    KossonSILVER0708 = Kanon {"Kanon 20. 149", "Folget dem, das euch offenbaret ist durch den JHWH und den Propheten, dann folget ihr den wahrheitlichen Gesetzen und Geboten und der siebenheitlichen Ordnung der Schöpfung."}
    KossonSILVER0709 = Kanon {"Kanon 20. 150", "Wendet euch ab von euren Göttern und Götzen, von euren Menschen, die ihr als Gott verehret und als Söhne und Töchter der Schöpfung, die ihr als Nebenbuhler an ihre Seite stellet, denn es gibt keine Schöpfung im Himmel (Universum) ausser der Schöpfung, die da ist die einzige und die eure."}
    KossonSILVER0710 = Kanon {"Kanon 20. 151", "Wendet euch ab von euren Götzendienern, die euch da irreführen und euch Falschheit und Lug und Trug predigen über selbsternannte oder durch sie erkorene Götter und Götzen."}
    KossonSILVER0711 = Kanon {"Kanon 20. 152", "Und wahrlich, wäre die Schöpfung nicht wahrliche Freiheit, und wären nicht frei im Willen der Entwicklungswege ihre Kreaturen, dann hätte sie selbst willentlich bestimmet, dass da Nebenbuhler neben ihr seien und Irrlehren und Verehrung und Anbetung, die da von ihr dann gefordert würden."}
    KossonSILVER0712 = Kanon {"Kanon 20. 153", "So dem aber nicht so ist und die Wege der Entwicklung dem Menschen freigesetzet sind, so ist der Mensch frei und wahrlich bestimmend und herrschend über sich selbst, und also ist er der Herr und Meister seiner selbst."}
    KossonSILVER0713 = Kanon {"Kanon 20. 154", "Es ermöglichet ihm dies, zu rechten wider das Gerechte und das Ungerechte und zu bestimmen allein seinen Weg des Gerechten oder Ungerechten."}
    KossonSILVER0714 = Kanon {"Kanon 20. 155", "Und da der Mensch lebet im freien Willen seiner Entscheidung, so suchet er nur den Vorteil durch geringsten Widerstand, der da lieget in der Abart des Mühelosen, wie die gedämmten Wasser des Teiches, die da ein Leck suchen, um ohne Widerstand zu entweichen, so sie als geballte Kraft alles zerstörend wegreissen und vernichten.»"}
    KossonSILVER0715 = Kanon {"Kanon 20. 156", "Und es spricht der JHWH zum Propheten: «Du bist nicht zum Hüter der Menschen gemachet, und also nicht zu einem Wächter über sie, und so auch nicht der JHWH."}
    KossonSILVER0716 = Kanon {"Kanon 20. 157", "Du aber bist gemachet aus eigener Kraft zum Propheten der Belehrung der Wahrheit unter den Menschen,so auch der JHWH ist Belehrender für den Propheten und für die Menschengeschlechter.»"}
    KossonSILVER0717 = Kanon {"Kanon 20. 158", "Und spreche belehrend zu den Menschen: «Schmähet nicht eure Freunde und schmähet nicht eure Feinde, und schmähet nicht die Gerechten und nicht die Ungerechten, denn Kreaturen der Schöpfung sind sie allesamt."}
    KossonSILVER0718 = Kanon {"Kanon 20. 159", "Bringet den Feinden Liebe entgegen, wie den Freunden, und bringet den Ungerechten Liebe entgegen, wie den Gerechten, doch seid wachsam und hütet euch vor Falschheit und Lug und Trug der Feinde und der Ungerechten, so ihr nicht erleidet Schaden durch sie."}
    KossonSILVER0719 = Kanon {"Kanon 20. 160", "Ist da ein Feind oder ein Ungerechter, so sind sie Mensch wie ihr, so ihr über sie und mit ihnen und für sie auch denken und handeln sollet als Mensch und in Liebe also, auch wenn euch geboten ist Wachsamkeit, dass ihr nicht Schaden leidet, wenn sie heimlich willens sind, euch Schaden und Falschheit und Lug und Trug zu bringen."}
    KossonSILVER0720 = Kanon {"Kanon 20. 161", "Seid freundlich und liebevoll und menschlich zu den Feinden und Ungerechten, und gebet ihnen Obdach, Speise und Trank und alles Notwendige, und entbietet ihnen gebührende Ehre und Respekt, denn sie sind Menschen als auch ihr Menschen seid, und alle seid ihr Kreaturen der gleichen Schöpfung, so die Gerechten als auch die Ungerechten."}
    KossonSILVER0721 = Kanon {"Kanon 20. 162", "Und schmähet keine Gefangenen, so ihr welche machet, wenn ihr von Feinden überfallen werdet und ihr euch erwehren müsset."}
    KossonSILVER0722 = Kanon {"Kanon 20. 163", "Ehret die Gefangenen, und bringet ihnen entgegen Freundlichkeit, Menschlichkeit, Liebe und Respekt, und gebet ihnen Nahrung und Trank und alles, was sie da brauchen."}
    KossonSILVER0723 = Kanon {"Kanon 20. 164", "Und schmähet da nicht die Alten und nicht die Jungen und nicht die Geschädigten am Körper, und schmähet nicht den Nächsten, der da eine andere Sprache spricht oder der da hat eine andere Hautfarbe, denn es ist jeder Mensch gleichwertig dem anderen, welcher Art und Rasse er auch immer ist, und ob er da ist alt oder jung, gelb, schwarz oder braun, oder weiss oder rot oder blau, und ob es da ist Kind, Weib oder Mann, und ob er da ist Feind oder Freund oder Gerechter, Ungerechter, Wissender oder Unwissender."}
    KossonSILVER0724 = Kanon {"Kanon 20. 165", "Und nicht ist ein Mensch einem anderen der Untertan, so also nicht als Leibeigener und nicht als Sklave, und nicht als Kämpfer eines Heeres den Führern, und also nicht als Volk der Obrigkeit und der Herrscher."}
    KossonSILVER0725 = Kanon {"Kanon 20. 166", "Und es sei auch nicht das Weib des Mannes Untertan, also auch nicht die Kinder seien Untertan der Eltern, und nicht der Mann der Untertan des Weibes, und nicht der Tagelöhner und Arbeiter seinem Meister, und es seien nicht die Eltern Untertan der Kinder, und also seien nicht die Menschen Untertan von Göttern und Götzen und von Verführern und von Priestern und ihren Sekten und Religionen."}
    KossonSILVER0726 = Kanon {"Kanon 20. 167", "Jeder Mensch sei gleichwertig dem Nächsten, und er stehe im Recht desgleichen zum Nächsten, so der Arme zum Reichen und in umgekehrter Form, und so jeglicher Mensch zum andern, was und wer und wie er auch immer sei, und es werde niemals in zweierlei Mass gemessen in irgendeiner Form zum Recht oder Unrecht des Menschen und der Lebensformen im gesamten überhaupt."}
    KossonSILVER0727 = Kanon {"Kanon 20. 168", "Und wahrlich, es teile der Besitzende sein Brot mit dem Besitzlosen in freiem Willen, und es sei nicht fordernd der Besitzlose oder der Besitzende beim Nächsten, sondern stets sei er gerecht und fragend, so er einer Hilfe oder einer Sache bedarf."}
    KossonSILVER0728 = Kanon {"Kanon 20. 169", "Es ist aber gegeben, dass ein jeder seine Pflicht tue in Erfüllung der Gesetze und Gebote und der siebenheitlichen Ordnung der Schöpfung, so also auch verrichte ein jeder seine ihm zugeordnete Arbeit, und ein jeder sei lehrsam und evolutiv und befolgend in der schöpferischen Ordnung, so nicht und niemals seien Schmarotzer und Parasiten, die da auf Hilfe hoffen und Hilfe fordern, deren sie nicht würdig sind und die da ihnen nicht kann gegeben sein.»"}
    KossonSILVER0729 = Kanon {"Kanon 20. 170", "Und spreche also belehrend zu den Menschen», sprichet der JHWH: «Lasset euch nicht irreleiten durch jene auf der Erde, die da nicht sind gehorchend den Gesetzen und Geboten der Schöpfung, denn sie leben willentlich den Weg der Mühelosigkeit und der Ungerechtigkeit, und sie sind stetig nur bringend falsche Gelüste und falsche Freuden und falsche Lehren, Gier, Unfrieden, Unliebe und Rache und Übel und alle Ungerechtigkeit der Welt bis zum Tode."}
    KossonSILVER0730 = Kanon {"Kanon 20. 171", "Sie sind Verlierer neben dem Weg der Schöpfung, und sie folgen nur einem Wahn, und sie lügen und trügen bloss."}
    KossonSILVER0731 = Kanon {"Kanon 20. 172", "Und wahrlich, viele führen andere irre, mit bösen Falschheiten, Lügen und Gelüsten, und sie tun dies im Mangel an Wissen oder wahrheitlich wider besseres Wissen, doch die ihnen hörig werden und falsche Wege gehen, die werden mitschuldig mit ihren Verführern."}
    KossonSILVER0732 = Kanon {"Kanon 20. 173", "Meidet offenkundige Ungerechtigkeit ebenso wie auch geheime Ungerechtigkeit, denn ob offenkundig oder geheim, beide sind sie von Falschheit, und wer da begehet Ungerechtigkeit in der einen oder anderen Form, der brichet das Recht, und er wird Lohn empfangen demgemäss."}
    KossonSILVER0733 = Kanon {"Kanon 20. 174", "Und lasset euch nicht irreführen durch Verführer des Glaubens, die da Götter und Götzen oder einen Gott neben die Schöpfung setzen, so ihr aber doch den Verführern hörig seid,so werdet ihr gefangen in deren Göttern und Götzen oder in deren Gott, und nur schwerlich oder überhaupt nicht werdet ihr hinfinden zur Schöpfung und zu ihrer Wahrheit."}
    KossonSILVER0734 = Kanon {"Kanon 20. 175", "Denn Götter und Götzen fordern Untertänigkeit und Demütigung vor ihnen, und sie schaffen Gewalt und Zwang, und also Ausbeutung und Versklavung durch Angst und Furcht, die da erzeugen Unfreiheit und Willensbruch im selbständigen und selbstentscheidenden Denken und Handeln."}
    KossonSILVER0735 = Kanon {"Kanon 20. 176", "So ihr also den Verführern gehorchet, so setzet ihr neben die Schöpfung Götter und Götzen, und ihr werdet unfrei und voller Angst und Furcht, und ihr seid gehorsam den Verführern und ihren Göttern und Götzen, wie geprügelte Hunde ihren Herren in hündischer Demut."}
    KossonSILVER0736 = Kanon {"Kanon 20. 177", "Denn es sind die Wege des Ungerechten ausgeleget mit allen Übeln der Welt, und da Verführer sind und Götter und Götzen, da herrschet Angst und Furcht und Schrecken vor Strafe, die Geltung erlangen soll, so den Verführern, Göttern und Götzen nicht gebuhlet werde."}
    KossonSILVER0737 = Kanon {"Kanon 20. 178", "Und es fordern Verführer, Ungerechte, Götter und Götzen grausame Strafen für Denkweisen und Handlungen wider ihre Falschheit und ihre falschen Gesetze und Gebote, und es werden dadurch die Gläubigen und Hörigen gezwänget in Angst und Furcht und Schrecken, so sie nicht mutig werden, sich zu widersetzen allen Falschheiten und den falschen Kräften, die da nur fordern Strafe und Tod."}
    KossonSILVER0738 = Kanon {"Kanon 20. 179", "Es ist ihnen mit ihren Gedanken verwehret der Weg zum Wissen und zur Wahrheit, und es ist dies getan durch die Angst und Furcht, durch die sie sind in Fesseln geschlagen und unfrei durch die Verführer und Götter und Götzen, die mit qualvollem und tödlichem Strafgericht drohen und mit höllischer Verdammnis, wenn nicht erfüllet werden ihre falschen Forderungen und Machtheischungen."}
    KossonSILVER0739 = Kanon {"Kanon 20. 180", "Und es ist den Gläubigen der Verführer und Götter und Götzen untergründig bewusst ihre Angst und Furcht, doch es sind die Drohungen für die höllischen Strafgerichte und Qualen, die sie da nicht freilassen."}
    KossonSILVER0740 = Kanon {"Kanon 20. 181", "Und die Drohenden sind die Verführer und ihre Götter und Götzen und die Sekten und Religionen und ihre falschen Gesetze und Gebote, und also sind es die Priester der Gläubigen, und also sind die Drohenden auch die Machtgierigen, die Gewinnsüchtigen und die Führenden von Kriegsheeren, und also sind es die Obrigkeiten des Volkes, die da in unwahrheitlicher Form und nach Gesetzen und Geboten des Menschen regieren, verführen, drangsalieren und morden, und die da nicht regieren und führen das Volk im Sinne des Geistes und der wahrlichen Freiheit, und nicht nach den Gesetzen und Geboten der Schöpfung."}
    KossonSILVER0741 = Kanon {"Kanon 20. 182", "Wahrlich, es vermag nicht ein Toter unter wahrheitlich Lebenden zu wandeln, so auch nicht ein Ungerechter Gerechtigkeit erlangen kann, ehe er sich nicht gewandelt hat zum Gerechten, so auch der Tote erst als neue Persönlichkeit wiedergeboren werden muss, ehe es ihm möglich sein kann, unter Lebenden zu wandeln."}
    KossonSILVER0742 = Kanon {"Kanon 20. 183", "Es mag so ein Mensch nur frei sein, wenn er da wandelt im strahlenden Lichte des Lebens und der Wahrheit, und es mag so nur jener Mensch unfrei sein und gefangen, der da stolpert und stürzet im Kerker der schwärzesten Finsternis."}
    KossonSILVER0743 = Kanon {"Kanon 20. 184", "Und die da, die in schwärzester Finsternis stolpern und stürzen, sie wähnen Ränke zu schmieden wider die Gerechten und wider die Schöpfung, doch sie schmieden nur Ränke wider sich selbst, allein, sie merken es nicht, denn sie sind dumm und blind, kleinmütig und schwach und willenlos gegen sich selbst."}
    KossonSILVER0744 = Kanon {"Kanon 20. 185", "Wahrlich, sie sind die Verlierer, und sie bringen sich selbst Erniedrigung und Strafe aus eigener Kraft."}
    KossonSILVER0745 = Kanon {"Kanon 20. 186", "Nicht wollen sie leben der Wahrheit der Schöpfung, die da selbst weder Strafe fordert noch Strafe erteilet, so der Mensch sich immerwährend selbst bestrafet für ungerechtes Tun und Handeln, wie auch das Kaninchen vom Adler geschlagen und getötet wird, wenn es den drohenden Schatten des herabstürzenden Adlers missachtet und nicht in die schützende Höhle fliehet."}
    KossonSILVER0746 = Kanon {"Kanon 20. 187", "Und so der Adler das durch eigene Kraft unvorsichtige Kaninchen schlaget und tötet, so ziehet der Mensch sich selbst Gericht und Strafe zu aus eigener Kraft, wenn in ihm lastet Unvorsichtigkeit und Wille zum falschen Denken und Handeln."}
    KossonSILVER0747 = Kanon {"Kanon 20. 188", "Darum: Wer sich gerecht leitet und ein schöpfungsgerechtes Leben führet, dem weitet sich die Brust und die Freude und die Liebe, das Wissen, die Wahrheit und die Weisheit,und er lasset sich nicht vom Feinde schlagen, denn er ist wachsam und vorsichtig."}
    KossonSILVER0748 = Kanon {"Kanon 20. 189", "Wer sich aber ungerecht leitet und ein schöpfungswidriges Leben führet, der enget sich die Brust, und seine Freude wird klein sein, und so auch seine Liebe, das Wissen und die Wahrheit und die Weisheit also."}
    KossonSILVER0749 = Kanon {"Kanon 20. 190", "Es ist das der Weg der Wahrheit und der Weg der Schöpfung, der da ist der wahrheitliche und der rechte, und wer ihn begehet, der lebet botmässig den Gesetzen und Geboten der Schöpfung."}
    KossonSILVER0750 = Kanon {"Kanon 20. 191", "Wahrlich, es sind verschiedenste Stufen für alle Menschen, je nach ihrem Tun, doch ist alles geregelt in den Gesetzen und Geboten und der siebenheitlichen Ordnung der Schöpfung, dass jeder belohnet werde nach seiner Stufe und nach seinem Tun."}
    KossonSILVER0751 = Kanon {"Kanon 20. 192", "Und es ist gegeben, dass allzeitlich jeder belohnet werde nach seiner Stufe und nach seinem Tun, in Gerechtigkeit und in Unparteilichkeit, so allzeitlich wahrlich jeglicher seinen Lohn erhaltet, wie er ihm gebühret und wie er bemessen ist nach seiner Lebensführung und nach seinen Taten, und so nach seinem Denken und Handeln."}
    KossonSILVER0752 = Kanon {"Kanon 20. 193", "Und es kann keiner vereiteln oder abweisen den Lohn, der ihm nach seiner Stufe zusteht, denn die schöpferische Gerechtigkeit kennet keinen Zwiespalt und weder Bevorzugung noch Benachteiligung."}
    KossonSILVER0753 = Kanon {"Kanon 20. 194", "Es handele daher der Mensch stetig nach seinem besten Vermögen, und also handele er in Erfüllung der Gesetze und Gebote und der siebenheitlichen Ordnung der Schöpfung, nach bestem Können und Vermögen, so er also recht tuet und stets gebet das Beste nach seinem Verstand und Können, und so er also wahrheitlich handelt und lebet und nicht ungerecht sei."}
    KossonSILVER0754 = Kanon {"Kanon 20. 195", "Und viele Irregeleitete, Unwissende und Ungerechte sind da, die da glauben, dass sie Opfer und Gebete (Betteleien und Flehungen usw.) darbringen könnten ihren Verführern und Göttern und Götzen und dass sie dadurch befreiet wären von der Schuld des Götzendienstes."}
    KossonSILVER0755 = Kanon {"Kanon 20. 196", "Aber wahrlich, ihre Opfer und Gebete (Betteleien und Flehungen usw.) sind an Götter und Götzen gerichtet und also nicht bestimmet für die Schöpfung."}
    KossonSILVER0756 = Kanon {"Kanon 20. 197", "Wer da unwissentlich bringet Opfer und Gebete (Betteleien und Flehungen) an Götter und Götzen und der da ist glaubend, dass er damit gelange an die Schöpfung, der ist nicht weniger und ist nicht mehr als einer der Ungerechten anderen, denn er frönet Götzen und Göttern und verleumdet die Wahrheit und die Schöpfung."}
    KossonSILVER0757 = Kanon {"Kanon 20. 198", "Und es wird sein den Ungerechten ihr eigenes Urteil und ihr selbstberufenes Gericht, und ihre selbsternannte Strafe also."}
    KossonSILVER0758 = Kanon {"Kanon 20. 199", "Und ihr Gericht und ihr Urteil und ihre Strafe werden nicht sein von der Schöpfung und nicht von ihren Gesetzen und Geboten,sondern sie werden sein erkoren durch die eigene Kraft und Macht des Menschen selbst, gemessen nach der Schwere der Tat des Vergehens und der Schuld."}
    KossonSILVER0759 = Kanon {"Kanon 20. 200", "Denn es ist der Mensch Herr über sich selbst, so er sein Leben bestimmet und darüber richtet in eigener Kraft und Verantwortung, so er sich auch selbst Strafe und Gericht auferleget oder von sich abwendet.»"}
    KossonSILVER0760 = Kanon {"Kanon 20. 201", "Und es sprichet der JHWH: «Menschengeschlechter, handelt nach eurem Vermögen und nach eurem Verstand und eurer Vernunft, und findet, wessen euer endgültiger Lohn sein wird für das Leben, das ihr führet."}
    KossonSILVER0761 = Kanon {"Kanon 20. 202", "Sehet aber, den Gerechten wird guter Lohn sein, doch den Ungerechten wird es nicht wohl ergehen."}
    KossonSILVER0762 = Kanon {"Kanon 20. 203", "Ihr opfert Früchte und Feldfrüchte und Getier und Weihrauch euren Göttern und Götzen, und ihr opfert Blut und gar Menschen und Dinge eures Reichtums, und ihr verbindet eure Opferungen mit plappernden Gebeten (Betteleien und Flehungen usw.), und ihr erhoffet durch sie von euren Göttern und Götzen Hilfe, Barmherzigkeit und Gnade."}
    KossonSILVER0763 = Kanon {"Kanon 20. 204", "Aber kein Gott wird euch zu irgendeiner Zeit Gehör, Hilfe, Barmherzigkeit oder Gnade zuteil werden lassen, und so wird auch nicht tun ein Götze oder sonstiges in gleicher oder ähnlicher Form, denn Götter sind Gewaltherrscher, die da nur Gewinn und Macht erstreben, und Götzen sind tote Kultdämonen ohne jegliche Kraft und Macht, und die da von Priestern und dergleichen benutzet werden zu eigensüchtigen und gewinnsüchtigen Zwecken."}
    KossonSILVER0764 = Kanon {"Kanon 20. 205", "Also opfert ihr eure Nahrung und euer Getier und euren Reichtum ohne Sinn und ohne begründbare Möglichkeit einer Hoffnung, und ihr folget damit nur einem wahngläubischen Kult, wobei ihr gar mordet und Menschenleben und anderes Leben vernichtet."}
    KossonSILVER0765 = Kanon {"Kanon 20. 206", "Sehet, Menschen der Erde, so ihr in Gerechtigkeit und Wahrheit lebet, so verehret (ehrwürdigt) ihr nur die Schöpfung, und also betet (ehrt) ihr nur zur Schöpfung, die da ist Geistform und so auch als winziges Teilstück in euch selbst, so ihr also auch beten (ehrwürdigen) könnet zum Teilstück Schöpfungsgeist in euch und zu euch selbst."}
    KossonSILVER0766 = Kanon {"Kanon 20. 207", "Und in keinster Weise ist die Schöpfung ähnlich in irgendwelchen Dingen den Göttern und Götzen, so sie auch niemals erliess Gesetze oder Gebote für kultische Handlungen für sich selbst, und also waren da niemals erschaffen von ihr Gesetze oder Gebote zur Opferdarbringung an sie in irgendwelcher Form."}
    KossonSILVER0767 = Kanon {"Kanon 20. 208", "Es ist die Schöpfung unbestechlich und nicht befangen in einer parteilichen Form, und so ist sie ohne Forderungen an ihre Kreationen, so sie nur ist Schöpfung und Erfüllung des Lebens und der Gesetze und Gebote und der siebenheitlichen Ordnung."}
    KossonSILVER0768 = Kanon {"Kanon 20. 209", "Also ist gegeben, dass sie dadurch nicht fordernd ist in irgendeiner Form und also nicht bedürftig nach irgendwelchen Gaben oder Opfern an sie durch ihre Kreationen."}
    KossonSILVER0769 = Kanon {"Kanon 20. 210", "Es bedarf die Schöpfung also keiner Opferdarbringungen, die da auch nicht wären entsprechend irgendeiner Ordnung von ihr, sondern die da nur sind Zerstörung von Leben, Nahrung und Gut in Sinnlosigkeit und Verschwendung durch kultische Ausartung des Menschen, und die da nur dienend sind als falsche Bezeugung und falsche Ehrlichkeit des Menschen, zum Zwecke demütiger Unterwerfung in Angst und Furcht, so also nicht dargebracht werden sollen Opferdarbringungen in irgendeiner Form."}
    KossonSILVER0770 = Kanon {"Kanon 20. 211", "Also stellet dar die Opferdarbringung einen Verstoss gegen die Gesetze und Gebote und gegen die siebenheitliche Ordnung der Schöpfung, die da nicht fordert oder erlaubet Zerstörung oder Mordung oder Verschwendung von irgendwelchem Leben oder von Gut oder Nahrung in welcher Form auch immer."}
    KossonSILVER0771 = Kanon {"Kanon 20. 212", "Und so die Schöpfung ist neutral in Ausgeglichenheit und in Liebe, und so sie ist unbestechlich und nicht befangen und ohne Forderung in irgendeiner Form gegenüber allem Leben und ihren Kreationen, so sind von ihr nur erstellet Weisungen und Gesetze und Gebote und die siebenheitliche Ordnung, die da beinhalten alle Richtlinien und Wegweiser der gerechten Lebensführung für jegliches Leben, das da entscheidet in eigenem freien Willen über den eigenen Weg des Lebens."}
    KossonSILVER0772 = Kanon {"Kanon 20. 213", "Und also dem so ist, ist auch nicht erstellet ein Tüpfelchen von einem Gesetz oder Gebot der Schöpfung, das da fordern würde von irgendeiner Lebensform, und so also auch nicht vom Menschen, dass ihr da entgegengebracht sein sollen Verehrung und Anbetung oder Gaben und Opferungen in irgendwelcher Form, denn nicht ist die Schöpfung fordernd in irgendwelchen Dingen, so also nicht in guter und nicht in böser Form."}
    KossonSILVER0773 = Kanon {"Kanon 20. 214", "Dem Menschen aber ist es freigestellet, und es gebühret ihm also, dass er entbiete der Schöpfung Ehrfurcht (Ehrung, Ehrerbietung, Ehrwürdigung, Ehrwürdigkeit) und Respekt und Ehrung, so auch gebühret jeglichem anderem Leben."}
    KossonSILVER0774 = Kanon {"Kanon 20. 215", "Und es lieget dem Menschen in seinem freien Willen, zu verrichten an die Schöpfung Ansprachen und Gebete (Ehrung)und Verehrung (Ehrwürdigung), so er sich konfrontiere mit der Kraft des wahrlichen Lebens in Kommunikation und also mit seinem Teilstück Schöpfungsgeist in sich selbst, der durch diesartige Gebete (Ehrungen) angesprochen und geformet wird in einsprechender Form und in wahrlichem Wissen, in Liebe, Weisheit und Wahrheit."}
    KossonSILVER0775 = Kanon {"Kanon 20. 216", "Also ergebet sich durch solcherlei Gebet (Ehrung) ein besseres Fühlen und Sicherheit im Innern des Menschen, so er Hoffnung heget und evolutionieret, Kraft dessen das durch das Gebet （Ehrung, Insichgehen, Selbsterkenntnis) angesprochene Teilstück Geist (Geistform) im Menschen für den Menschen selbst unbewusst seine Kräfte entfaltet und evolutive Kraft und Impulse spendet, wodurch erfüllet werden die Wünsche und Hoffnungen des Menschen."}
    KossonSILVER0776 = Kanon {"Kanon 20. 217", "Also ist ein Gebet (Ehrung, Insichgehen, Selbsterkenntnis) gerichtet an die Kraft Schöpfung, oder an das Teilstück Geist im Menschen, oder an das materielle Bewusstsein und das materielle Unterbewusstsein selbst, so es wirksame Kräfte im Menschen selbst anrege und zur Wirkung bringe."}
    KossonSILVER0777 = Kanon {"Kanon 20. 218", "Ob ein Gebet (Ehrung, ein Insichgehen, eine Selbsterkenntnis) gerichtet ist an die Schöpfung oder an die materiellen Bewusstseinsformen selbst, es wird angesprochen und angereget dadurch in direkter Form das Teilstück Schöpfungsgeist im Menschen, das da wirksam wird durch die gute Hoffnung und das Wissen des Menschen, dass der Geist durch seine Kraft wahrheitlich Hilfe erteilet."}
    KossonSILVER0778 = Kanon {"Kanon 20. 219", "Also ist ein Gebet (eine Ehrung, ein Insichgehen, eine Selbsterkenntnis) von gleichem Wert, ob dieses gerichtet wird direkt an die Schöpfung oder an die materiellen Bewusstseinsformen oder an das Teilstück Geist des Menschen, der da ja ist ein winziges Teilstück der Schöpfung selbst, so er also verfüget auch über die Kraft der Geistesenergie der Schöpfung, die da ist neutral in allen Dingen und ausgeglichen also, und die niemals in eine andere Form verändert oder gewandelt werden kann, weder die Geistesenergie der Schöpfung selbst noch das Teilstück Geist im Menschen."}
    KossonSILVER0779 = Kanon {"Kanon 20. 220", "Also ist gesaget, dass die Schöpfung in keiner Form fordernd oder befehlend ist, dass also zu ihr gebetet (gebettelt und gefleht) werden müsste, denn eine Anbetung des Schöpfungsgeistes im ganzen oder das Teilstück Geist im Menschen, das ein Teil der ganzen Schöpfung ja ist, das lieget allein im Willen und Ermessen des Menschen, so also auch in seiner Vernunft und in seinem Verstand, die da wissend sein sollen, dass Gebete (Ehrungen und Bestimmungen) ein gutes und wirksames Mittel der Hilfe sind für den Menschen und für seine Evolution und so auch Hoffnung, Wunscherfüllung, Sicherheit und Geborgenheit."}
    KossonSILVER0780 = Kanon {"Kanon 20. 221", "Wahrlich nämlich geben Gebete (Ehrungen, Ehrwürdigungen und Bestimmungen) dem Menschen Sicherheit des Wissens und der Wahrheit, so er in wahrlicher Wahrheit wandelt und seine Gebete (Ehrungen, Ehrwürdigungen und Bestimmungen) verrichtet in wahrlicher Form."}
    KossonSILVER0781 = Kanon {"Kanon 20. 222", "Es sind Gebete (Erwürdigungen) der Anbetung (Ehrwürdigkeit) an die Schöpfung, an die materiellen Bewusstseinsformen und an den Geist im Menschen nicht und niemals zu verrichten in einer Form von Demut oder Bangen oder Angst und Furcht, denn so dem fälschlich getan wird, sinket die Moral des Menschen und also die Sicherheit und das Vertrauen um die Kraft der Geistenergie."}
    KossonSILVER0782 = Kanon {"Kanon 20. 223", "Es fordern weder die Schöpfung noch der Geistteil im Menschen Demut, Unterwürfigkeit oder Bangen und Angst und Furcht, denn es soll der Mensch frei sein in jeder Form, so er nicht zu bangen und nicht zu fürchten brauchet."}
    KossonSILVER0783 = Kanon {"Kanon 20. 224", "Und erfolget Anbetung an die Schöpfung, an die materiellen Bewusstseinsformen oder an den Geist im Menschen, dann darf die Anbetung nicht sein derart, dass da wären Huldigung, Anflehung oder Verherrlichung,sondern es soll sein nur Lobpreisung (hohe Ehrung) in Ehrfurcht (Ehrung, Ehrerbietung, Ehrwürdigung, Ehrwürdigkeit) und Respekt, Andacht und Ansuchung in wahrlicher Form."}
    KossonSILVER0784 = Kanon {"Kanon 20. 225", "Und so ihr euer Gebet (Ehrung, Insichgehen, Selbsterkenntnis) verrichtet, so sollet ihr nicht plappern wie die Gläubigen der Sekten und Religionen und wie die Heiden und sonstigen Gottdiener und Götzendiener, die da Anbetungen (Betteleien und Flehungen) in Demut verrichten, mit Gebeten (Plappereien), die da sind endlose Wiederholungen und auf Namen von Götzen und Göttern bezogen in unvernünftiger und folgefalscher Form."}
    KossonSILVER0785 = Kanon {"Kanon 20. 226", "Verrichtet ihr aber als Gerechte und Wissende eure Gebete (Ehrungen), dann sollen sie sein klar und sinnvoll und folgerichtig im Wort und in der Wahrheit."}
    KossonSILVER0786 = Kanon {"Kanon 20. 227", "Und so ihr betet (ehrt und bestimmt), so verrichtet euer Gebet (Ehre und Bestimmung) nach freiem Willen und nach freiem Sinn, so ihr selbst Worte zu vernünftigen und logischen Gebeten （Ehrungen und Bestimmungen) bildet, als ihr aber auch gebrauchet und sprechet Gebete （Ehrungen, Würdigungen, ein Insichgehen und Selbsterkenntnis), die da sind zweckbestimmend in allgemeiner Form und bestimmet zur Evolution und zur gerechten Lebensführung."}
    KossonSILVER0787 = Kanon {"Kanon 20. 228", "Habet acht bei euren Gebeten (Ehrungen und Bestimmungen), dass ihr da betet (ehrt und bestimmt) in Frömmigkeit, was da ist gerechtes und wahrheitliches Beten (Ehren, Würdigen usw.) in Wissen und Liebe, so ihr also sollet wählen die richtigen Worte und die richtige Form der Gebetsdarbringung (Ehre- und Bestimmungsdarbringung)."}
    KossonSILVER0788 = Kanon {"Kanon 20. 229", "Und wenn ihr betet (ehrt und bestimmt), sollet ihr sein achtsam darauf, dass ihr betet (ehrt und bestimmt) in Ehrlichkeit und nicht wie die Heuchler, die da sind Ungerechte und die da gerne stehen und beten (heucheln) an den Ecken der Häuser und auf den Strassen und Gassen, so sie damit scheinen wollen vor ihren Mitmenschen in Selbstsucht und Ansehen, und so sie damit neue Anhänger gewinnen und Profit ertrügen."}
    KossonSILVER0789 = Kanon {"Kanon 20. 230", "Und wahrlich, wenn ihr betet (ehrt und bestimmt), dann sollet ihr die Allmacht der Schöpfung anrufen, die da ist die Allmacht des Schöpfungsgeistes und des Geistes und der materiellen Bewusstseinsformen im Menschen."}
    KossonSILVER0790 = Kanon {"Kanon 20. 231", "Und ihr sollet im Gebet (Ehrung, Insichgehen, Selbsterkenntnis) nicht selbstsüchtig sein und nicht unverständig und nicht untertänig wie die Unwissenden, die Heuchler und Ungerechten, die da sinnen, dass sie dadurch von ihren falschen Göttern und Götzen erhöret würden, wenn sie viele Worte und unsinnig sprechen und flehen in Demut oder Befehl."}
    KossonSILVER0791 = Kanon {"Kanon 20. 232", "Wahrlich aber, der Schöpfungsgeist bedarf nicht vieler Worte, und so also auch nicht der Geist und die materiellen Bewusstseinsformen im Menschen; des Wissens darum bedarf es aber, dass der Geist machtvoll ist und gross und unendlich in seiner Kraft, durch die er wahrlich zu versetzen vermag Berge von einem Ort zum anderen Ort."}
    KossonSILVER0792 = Kanon {"Kanon 20. 233", "Verstehet ihr aber nicht direkt zur Allmacht des Geistes zu beten (zu ehren und zu bestimmen), dann behelfet euch mit einer Sache, durch die ihr euch vorbereiten und konzentrieren könnet, so ihr in Konzentration den Weg für eure Gedanken zur Allmacht des Geistes (Bewusstseins) findet."}
    KossonSILVER0793 = Kanon {"Kanon 20. 234", "Verstehet aber die Sache des helfenden Mittels in richtiger Form, denn wahrlich ist es nur ein Gerät der Hilfe zur Konzentration, das in sich selbst nicht irgendwelche Kraft besitzet und auch kein Mittler ist zwischen euch und der Allmacht des Geistes, denn wahrlich bleibet das Hilfsgerät nur ein Hilfsgerät allemal."}
    KossonSILVER0794 = Kanon {"Kanon 20. 235", "Gleichet also nie den Priestern der Kulte und Sekten und Religionen, und gleichet nie den Ungerechten,den Heuchlern, den Unverständigen und den Unwissenden und Selbstsüchtigen, die da Hilfsmittel zu Kultgegenständen erheben und zu Reliquien und die da ihre Kulthilfsmittel verehren in falschem Glauben, dass die Allmacht des Geistes in ihnen wohne."}
    KossonSILVER0795 = Kanon {"Kanon 20. 236", "Seid wissend also, dass auch durch die Nutzung von Hilfsmitteln und Heiligtümern (ehrwürdiglose Dinge,Gegenstände usw.)die Allmacht des Geistes stets in euch selbst wohnet, wie es auch gegeben ist in der Schöpfung selbst, die da ist Geist selbst, und der da ist also in allem Lebendigen, das da ist geschaffen von der Schöpfung."}
    KossonSILVER0796 = Kanon {"Kanon 20. 237", "Verrichtet daher eure Gebete (Ehrungen und Bestimmungen)als Wissende und Gerechte, und darum sollet ihr beten (ehren und bestimmen) folgerichtig in eigenen Gebeten (Ehrungen und Bestimmungen) oder in Gebeten (Ehrungen und Bestimmungen) von allgemeinem Wert und also:"}
    KossonSILVER0797 = Kanon {"Kanon 20. 238", "Mein Geist (Bewusstsein), der (das) du bist in Allmacht."}
    KossonSILVER0798 = Kanon {"Kanon 20. 239", "Dein Name sei geheiligt (kontrolliert, geehrwürdigt)."}
    KossonSILVER0799 = Kanon {"Kanon 20. 240", "Dein Reich offenbare und inkarniere sich in mir."}
    KossonSILVER0800 = Kanon {"Kanon 20. 241", "Deine Kraft entfalte sich in mir, und sie sei wirksam im Irdischen und im Geistigen in ausgeglichener und evolutiver Form."}
    KossonSILVER0801 = Kanon {"Kanon 20. 242", "Mein tägliches Brot gib mir heute, wessen ich da bedarf im Irdischen und im Geistigen."}
    KossonSILVER0802 = Kanon {"Kanon 20. 243", "Und lasse mich erkennen und beheben meine Fehler, und also lasse mich erkennen und erlernen die Wahrheit und die Liebe und die Weisheit und das Wissen."}
    KossonSILVER0803 = Kanon {"Kanon 20. 244", "Erfülle mich mit Kraft gegen Versuchung und Verwirrung, und also gib mir Kraft, so ich mich durch sie vom Irrtum erlöse."}
    KossonSILVER0804 = Kanon {"Kanon 20. 245", "Ich lebe nach deinen Gesetzen und nach deinen Geboten, denn dein ist das Reich des Geistes und des Lebens, und die Kraft und das Wissen, und die Wahrheit und Weisheit und Liebe in Allzeitlichkeit. SO SEI ES (Amen)."}
    KossonSILVER0805 = Kanon {"Kanon 20. 246", "Und wenn ihr eure Gebete (Ehrungen und Bestimmungen) verrichtet an den Geist der Schöpfung oder an das Teilstück Schöpfungsgeist oder an die materiellen Bewusstseinsformen in euch selbst, dann betet (ehrt und bestimmt) ihr zu nur einer Kraft, denn der Schöpfungsgeist ist sowohl in der Schöpfung selbst als auch in allen Schöpfungen der Schöpfung, so auch im Menschen also."}
    KossonSILVER0806 = Kanon {"Kanon 20. 247", "Und wenn ihr Gebete (Ehrungen und Bestimmungen) verrichtet an den Schöpfungsgeist selbst oder an den Geist in euch oder an die materiellen Bewusstseinsformen in euch, dann wird euch gegeben, wessen ihr bedürfet und um wessen ihr betet (ehrt und bestimmt), doch es ist gegeben, dass ihr seid wahrheitlich im Gebet (Ehrung, Würdigung, Insichgehen, Selbsterkenntnis) und im Wissen darum, dann werdet ihr empfangen."}
    KossonSILVER0807 = Kanon {"Kanon 20. 248", "Wenn ihr aber nicht seid Wissende, sondern Gläubige von irren Lehren, dass da die Kraft und der Geist nicht in euch selbst wohnend seien und dass da die Schöpfung nicht sei wahrheitlich und nicht wohnend in allem Lebendigen und von gleicher Art, wie das Teilstück Schöpfungsgeistform im Menschen selbst, dann seid ihr ohne Wissen und frönet fälschlich Göttern und Götzen oder einer sonstigen schöpfungswidrigen Glaubensdienerei anderer Form, so ihr lebet in geistunwirksamer Armut und in wissensloser Unzulänglichkeit."}
    KossonSILVER0808 = Kanon {"Kanon 20. 249", "Es ist aber auch gegeben den in falschem Glauben Handelnden und Lebenden, dass sie hie und da empfangen, was sie in irrem Glauben von zweckentfremdeten Hilfsmitteln, Reliquien und Heiligtümern (ehrwürdiglose Dinge, Gegenstände usw.) und von Göttern und Götzen erbitten und erflehen, doch sie empfangen dabei nur aus starrem und irrem Glauben heraus, und so also in gesetzwidriger Form der Zwangssuggestion, und also ohne Wissen um die wirkliche Wahrheit."}
    KossonSILVER0809 = Kanon {"Kanon 20. 250", "Wahrlich, selig sind nur die, die der wahrlichen Wahrheit und dem wahrlichen Wissen ihr Leben entbieten, denn sie allein empfangen in Ehrlichkeit und in Erfüllung der Gesetze und Gebote der Schöpfung und in wahrheitlicher Liebe und in Weisheit."}
    KossonSILVER0810 = Kanon {"Kanon 20. 251", "Die Schöpfung ist Ehrlichkeit und Erfüllung, und ihr gemäss geartet ist alles Leben, so der Mensch also, der zu erfüllen hat in eigener Freiheit und Ehrlichkeit nur seine in ihm schlummernden Veranlagungen, die da gegeben sind durch die Gesetze und Gebote der Schöpfung."}
    KossonSILVER0811 = Kanon {"Kanon 20. 252", "Also ist der Mensch im Innersten der Schöpfung gleich, so nämlich im Geiste, und er mache sich auch äusserlich seinem Geiste gleich, also er erfüllen möge auch im Äusseren die Gesetze und Gebote der Schöpfung.»"}

    //Kanon 21
    KossonSILVER0812 = Kanon {"Kanon 21.   1", ""}
    KossonSILVER0813 = Kanon {"Kanon 21.   2", ""}
    KossonSILVER0814 = Kanon {"Kanon 21.   3", ""}
    KossonSILVER0815 = Kanon {"Kanon 21.   4", ""}
    KossonSILVER0816 = Kanon {"Kanon 21.   5", ""}
    KossonSILVER0817 = Kanon {"Kanon 21.   6", ""}
    KossonSILVER0818 = Kanon {"Kanon 21.   7", ""}
    KossonSILVER0819 = Kanon {"Kanon 21.   8", ""}
    KossonSILVER0820 = Kanon {"Kanon 21.   9", ""}
    KossonSILVER0821 = Kanon {"Kanon 21.  10", ""}
    KossonSILVER0822 = Kanon {"Kanon 21.  11", ""}
    KossonSILVER0823 = Kanon {"Kanon 21.  12", ""}
    KossonSILVER0824 = Kanon {"Kanon 21.  13", ""}
    KossonSILVER0825 = Kanon {"Kanon 21.  14", ""}
    KossonSILVER0826 = Kanon {"Kanon 21.  15", ""}
    KossonSILVER0827 = Kanon {"Kanon 21.  16", ""}
    KossonSILVER0828 = Kanon {"Kanon 21.  17", ""}
    KossonSILVER0829 = Kanon {"Kanon 21.  18", ""}
    KossonSILVER0830 = Kanon {"Kanon 21.  19", ""}
    KossonSILVER0831 = Kanon {"Kanon 21.  20", ""}
    KossonSILVER0832 = Kanon {"Kanon 21.  21", ""}

    //Kanon 22
    KossonSILVER0833 = Kanon {"Kanon 22.   1", ""}
    KossonSILVER0834 = Kanon {"Kanon 22.   2", ""}
    KossonSILVER0835 = Kanon {"Kanon 22.   3", ""}
    KossonSILVER0836 = Kanon {"Kanon 22.   4", ""}
    KossonSILVER0837 = Kanon {"Kanon 22.   5", ""}
    KossonSILVER0838 = Kanon {"Kanon 22.   6", ""}
    KossonSILVER0839 = Kanon {"Kanon 22.   7", ""}
    KossonSILVER0840 = Kanon {"Kanon 22.   8", ""}
    KossonSILVER0841 = Kanon {"Kanon 22.   9", ""}
    KossonSILVER0842 = Kanon {"Kanon 22.  10", ""}
    KossonSILVER0843 = Kanon {"Kanon 22.  11", ""}
    KossonSILVER0844 = Kanon {"Kanon 22.  12", ""}
    KossonSILVER0845 = Kanon {"Kanon 22.  13", ""}
    KossonSILVER0846 = Kanon {"Kanon 22.  14", ""}
    KossonSILVER0847 = Kanon {"Kanon 22.  15", ""}
    KossonSILVER0848 = Kanon {"Kanon 22.  16", ""}
    KossonSILVER0849 = Kanon {"Kanon 22.  17", ""}
    KossonSILVER0850 = Kanon {"Kanon 22.  18", ""}
    KossonSILVER0851 = Kanon {"Kanon 22.  19", ""}
    KossonSILVER0852 = Kanon {"Kanon 22.  20", ""}
    KossonSILVER0853 = Kanon {"Kanon 22.  21", ""}
    KossonSILVER0854 = Kanon {"Kanon 22.  22", ""}
    KossonSILVER0855 = Kanon {"Kanon 22.  23", ""}
    KossonSILVER0856 = Kanon {"Kanon 22.  24", ""}
    KossonSILVER0857 = Kanon {"Kanon 22.  25", ""}
    KossonSILVER0858 = Kanon {"Kanon 22.  26", ""}
    KossonSILVER0859 = Kanon {"Kanon 22.  27", ""}
    KossonSILVER0860 = Kanon {"Kanon 22.  28", ""}
    KossonSILVER0861 = Kanon {"Kanon 22.  29", ""}
    KossonSILVER0862 = Kanon {"Kanon 22.  30", ""}
    KossonSILVER0863 = Kanon {"Kanon 22.  31", ""}
    KossonSILVER0864 = Kanon {"Kanon 22.  32", ""}
    KossonSILVER0865 = Kanon {"Kanon 22.  33", ""}
    KossonSILVER0866 = Kanon {"Kanon 22.  34", ""}
    KossonSILVER0867 = Kanon {"Kanon 22.  35", ""}
    KossonSILVER0868 = Kanon {"Kanon 22.  36", ""}
    KossonSILVER0869 = Kanon {"Kanon 22.  37", ""}
    KossonSILVER0870 = Kanon {"Kanon 22.  38", ""}
    KossonSILVER0871 = Kanon {"Kanon 22.  39", ""}
    KossonSILVER0872 = Kanon {"Kanon 22.  40", ""}
    KossonSILVER0873 = Kanon {"Kanon 22.  41", ""}
    KossonSILVER0874 = Kanon {"Kanon 22.  42", ""}
    KossonSILVER0875 = Kanon {"Kanon 22.  43", ""}
    KossonSILVER0876 = Kanon {"Kanon 22.  44", ""}
    KossonSILVER0877 = Kanon {"Kanon 22.  45", ""}
    KossonSILVER0878 = Kanon {"Kanon 22.  46", ""}
    KossonSILVER0879 = Kanon {"Kanon 22.  47", ""}
    KossonSILVER0880 = Kanon {"Kanon 22.  48", ""}
    KossonSILVER0881 = Kanon {"Kanon 22.  49", ""}
    KossonSILVER0882 = Kanon {"Kanon 22.  50", ""}
    KossonSILVER0883 = Kanon {"Kanon 22.  51", ""}
    KossonSILVER0884 = Kanon {"Kanon 22.  52", ""}
    KossonSILVER0885 = Kanon {"Kanon 22.  53", ""}
    KossonSILVER0886 = Kanon {"Kanon 22.  54", ""}
    KossonSILVER0887 = Kanon {"Kanon 22.  55", ""}
    KossonSILVER0888 = Kanon {"Kanon 22.  56", ""}
    KossonSILVER0889 = Kanon {"Kanon 22.  57", ""}
    KossonSILVER0890 = Kanon {"Kanon 22.  58", ""}
    KossonSILVER0891 = Kanon {"Kanon 22.  59", ""}
    KossonSILVER0892 = Kanon {"Kanon 22.  60", ""}
    KossonSILVER0893 = Kanon {"Kanon 22.  61", ""}
    KossonSILVER0894 = Kanon {"Kanon 22.  62", ""}
    KossonSILVER0895 = Kanon {"Kanon 22.  63", ""}
    KossonSILVER0896 = Kanon {"Kanon 22.  64", ""}
    KossonSILVER0897 = Kanon {"Kanon 22.  65", ""}
    KossonSILVER0898 = Kanon {"Kanon 22.  66", ""}
    KossonSILVER0899 = Kanon {"Kanon 22.  67", ""}
    KossonSILVER0900 = Kanon {"Kanon 22.  68", ""}
    KossonSILVER0901 = Kanon {"Kanon 22.  69", ""}
    KossonSILVER0902 = Kanon {"Kanon 22.  70", ""}

    //Kanon 23
    KossonSILVER0903 = Kanon {"Kanon 23.   1", ""}
    KossonSILVER0904 = Kanon {"Kanon 23.   2", ""}
    KossonSILVER0905 = Kanon {"Kanon 23.   3", ""}
    KossonSILVER0906 = Kanon {"Kanon 23.   4", ""}
    KossonSILVER0907 = Kanon {"Kanon 23.   5", ""}
    KossonSILVER0908 = Kanon {"Kanon 23.   6", ""}
    KossonSILVER0909 = Kanon {"Kanon 23.   7", ""}
    KossonSILVER0910 = Kanon {"Kanon 23.   8", ""}
    KossonSILVER0911 = Kanon {"Kanon 23.   9", ""}
    KossonSILVER0912 = Kanon {"Kanon 23.  10", ""}
    KossonSILVER0913 = Kanon {"Kanon 23.  11", ""}
    KossonSILVER0914 = Kanon {"Kanon 23.  12", ""}
    KossonSILVER0915 = Kanon {"Kanon 23.  13", ""}
    KossonSILVER0916 = Kanon {"Kanon 23.  14", ""}
    KossonSILVER0917 = Kanon {"Kanon 23.  15", ""}
    KossonSILVER0918 = Kanon {"Kanon 23.  16", ""}
    KossonSILVER0919 = Kanon {"Kanon 23.  17", ""}
    KossonSILVER0920 = Kanon {"Kanon 23.  18", ""}
    KossonSILVER0921 = Kanon {"Kanon 23.  19", ""}
    KossonSILVER0922 = Kanon {"Kanon 23.  20", ""}
    KossonSILVER0923 = Kanon {"Kanon 23.  21", ""}
    KossonSILVER0924 = Kanon {"Kanon 23.  22", ""}
    KossonSILVER0925 = Kanon {"Kanon 23.  23", ""}
    KossonSILVER0926 = Kanon {"Kanon 23.  24", ""}
    KossonSILVER0927 = Kanon {"Kanon 23.  25", ""}
    KossonSILVER0928 = Kanon {"Kanon 23.  26", ""}
    KossonSILVER0929 = Kanon {"Kanon 23.  27", ""}
    KossonSILVER0930 = Kanon {"Kanon 23.  28", ""}

    //Kanon 25
    KossonSILVER0931 = Kanon {"Kanon 25.   1", ""}
    KossonSILVER0932 = Kanon {"Kanon 25.   2", ""}
    KossonSILVER0933 = Kanon {"Kanon 25.   3", ""}
    KossonSILVER0934 = Kanon {"Kanon 25.   4", ""}
    KossonSILVER0935 = Kanon {"Kanon 25.   5", ""}
    KossonSILVER0936 = Kanon {"Kanon 25.   6", ""}
    KossonSILVER0937 = Kanon {"Kanon 25.   7", ""}
    KossonSILVER0938 = Kanon {"Kanon 25.   8", ""}
    KossonSILVER0939 = Kanon {"Kanon 25.   9", ""}
    KossonSILVER0940 = Kanon {"Kanon 25.  10", ""}
    KossonSILVER0941 = Kanon {"Kanon 25.  11", ""}
    KossonSILVER0942 = Kanon {"Kanon 25.  12", ""}
    KossonSILVER0943 = Kanon {"Kanon 25.  13", ""}
    KossonSILVER0944 = Kanon {"Kanon 25.  14", ""}
    KossonSILVER0945 = Kanon {"Kanon 25.  15", ""}
    KossonSILVER0946 = Kanon {"Kanon 25.  16", ""}
    KossonSILVER0947 = Kanon {"Kanon 25.  17", ""}
    KossonSILVER0948 = Kanon {"Kanon 25.  18", ""}
    KossonSILVER0949 = Kanon {"Kanon 25.  19", ""}
    KossonSILVER0950 = Kanon {"Kanon 25.  20", ""}
    KossonSILVER0951 = Kanon {"Kanon 25.  21", ""}
    KossonSILVER0952 = Kanon {"Kanon 25.  22", ""}
    KossonSILVER0953 = Kanon {"Kanon 25.  23", ""}
    KossonSILVER0954 = Kanon {"Kanon 25.  24", ""}
    KossonSILVER0955 = Kanon {"Kanon 25.  25", ""}
    KossonSILVER0956 = Kanon {"Kanon 25.  26", ""}
    KossonSILVER0957 = Kanon {"Kanon 25.  27", ""}
    KossonSILVER0958 = Kanon {"Kanon 25.  28", ""}
    KossonSILVER0959 = Kanon {"Kanon 25.  29", ""}
    KossonSILVER0960 = Kanon {"Kanon 25.  30", ""}
    KossonSILVER0961 = Kanon {"Kanon 25.  31", ""}
    KossonSILVER0962 = Kanon {"Kanon 25.  32", ""}
    KossonSILVER0963 = Kanon {"Kanon 25.  33", ""}
    KossonSILVER0964 = Kanon {"Kanon 25.  34", ""}
    KossonSILVER0965 = Kanon {"Kanon 25.  35", ""}
    KossonSILVER0966 = Kanon {"Kanon 25.  36", ""}
    KossonSILVER0967 = Kanon {"Kanon 25.  37", ""}
    KossonSILVER0968 = Kanon {"Kanon 25.  38", ""}
    KossonSILVER0969 = Kanon {"Kanon 25.  39", ""}
    KossonSILVER0970 = Kanon {"Kanon 25.  40", ""}
    KossonSILVER0971 = Kanon {"Kanon 25.  41", ""}
    KossonSILVER0972 = Kanon {"Kanon 25.  42", ""}
    KossonSILVER0973 = Kanon {"Kanon 25.  43", ""}
    KossonSILVER0974 = Kanon {"Kanon 25.  44", ""}
    KossonSILVER0975 = Kanon {"Kanon 25.  45", ""}
    KossonSILVER0976 = Kanon {"Kanon 25.  46", ""}
    KossonSILVER0977 = Kanon {"Kanon 25.  47", ""}
    KossonSILVER0978 = Kanon {"Kanon 25.  48", ""}
    KossonSILVER0979 = Kanon {"Kanon 25.  49", ""}
    KossonSILVER0980 = Kanon {"Kanon 25.  50", ""}
    KossonSILVER0981 = Kanon {"Kanon 25.  51", ""}
    KossonSILVER0982 = Kanon {"Kanon 25.  52", ""}
    KossonSILVER0983 = Kanon {"Kanon 25.  53", ""}
    KossonSILVER0984 = Kanon {"Kanon 25.  54", ""}
    KossonSILVER0985 = Kanon {"Kanon 25.  55", ""}
    KossonSILVER0986 = Kanon {"Kanon 25.  56", ""}
    KossonSILVER0987 = Kanon {"Kanon 25.  57", ""}
    KossonSILVER0988 = Kanon {"Kanon 25.  58", ""}
    KossonSILVER0989 = Kanon {"Kanon 25.  59", ""}
    KossonSILVER0990 = Kanon {"Kanon 25.  60", ""}
    KossonSILVER0991 = Kanon {"Kanon 25.  61", ""}
    KossonSILVER0992 = Kanon {"Kanon 25.  62", ""}
    KossonSILVER0993 = Kanon {"Kanon 25.  63", ""}
    KossonSILVER0994 = Kanon {"Kanon 25.  64", ""}
    KossonSILVER0995 = Kanon {"Kanon 25.  65", ""}
    KossonSILVER0996 = Kanon {"Kanon 25.  66", ""}
    KossonSILVER0997 = Kanon {"Kanon 25.  67", ""}
    KossonSILVER0998 = Kanon {"Kanon 25.  68", ""}
    KossonSILVER0999 = Kanon {"Kanon 25.  69", ""}
    KossonSILVER1000 = Kanon {"Kanon 25.  70", ""}
)

var (
    //Kanon 31
    KossonCOPPER1206 = Kanon {"Kanon 31. 470","Lügenhaftigkeit, Tollkünheit, Falscheit, Torheit, Narrheit, Blödheit, Habsucht, Gier, Laster, Zuchtlosigkeit, Unreinlichkeit, Grausamkeit, Machtsucht, Herrschsucht, Gewalt, Unzucht, Ausartung, Selbstsucht, Geiz, Faulheit, Rache, Hass, Freudlosigkeit, Angst und Furcht, Töten ohne Notwehr, Zwietracht und Streit und alle Übel sind das Unglück und das Leben des Täglichen aller jener, die da sind Gläubige von Sekten und Religionen, denn sie sind Unwissende und Ungerechte, die da in Irrlehren ein falsches Leben führen."}
    KossonCOPPER1344 = Kanon {"Kanon 31. 608", "Es sei der Mensch im Essen und im Tranke auch gemässiget, so er sich nicht hingebe an Frass und Völlerei, dergleichen aber, wird es doch getan, kommet gleich dem Verfälschen von Essbarem, also beides fallet in Unverstand und Schwachheit und also im schwaches Sinnen und Blödheit, denn beides führet zu Krankheit und Übel im Körper und im Denken und Fühlen, gleichsam wenn da vom Menschen Essbares vom Getier geraubet und gegessen wird, weshalben viele Menschein leiden an mannigfachen Krankheiten und an schwachem und langsamem Denken und an Schwachheit aller Art."}
    KossonCOPPER1898 = Kanon {"Kanon 31.1162","Semjasa"}
    KossonCOPPER1899 = Kanon {"Kanon 31.1163","Asasel"}
    KossonCOPPER1900 = Kanon {"Kanon 31.1164","Andanj"}
    KossonCOPPER1901 = Kanon {"Kanon 31.1165","Lunera"}
    KossonCOPPER1902 = Kanon {"Kanon 31.1166","Ezekeel"}
    KossonCOPPER1903 = Kanon {"Kanon 31.1167","Daanel"}
    KossonCOPPER1904 = Kanon {"Kanon 31.1168","Urakjbaramel"}

    //Kanon 32
    KossonCOPPER5548 = Kanon {"Kanon 32.2412","Schimpfnamen tun nicht weh."}

    //Kanon 49
    KossonCOPPER8792 = Kanon {"Kanon 49.  91","Der natürliche Menschenverstand wehrt sich zwar dagegen, den Kreis ungeeigneter irrealer religiöser Dogmen in das Quadrat der fordernden Wahrheit hineinpressen zu lassen, die ganz besonders durch die Neuzeit der Wassermann-Aera durch neue Religionsfanatiker, Sektierer und Betrüger zur blühenden Existenz gebracht werden; doch aber ist zur gegenwärtigen Zeit der Religionswahnsinn noch zu stark verbreitet, um eine nützliche Abwehr in grossem Masse erstellen zu können, folglich gegen die Jahrtausendwende viele Tausende von Menschen in sektiererischem Wahn durch Mord und Selbstmord sinnlos sterben werden."}
)