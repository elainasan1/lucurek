package commands

import (
	"filec2/Handlers/sessions"
)

func init() {
	Register(&Command{
		Name: "tokyo",
		Execute: func(session *sessions.Session, cmd []string) error {

			if err := session.Clear(); err != nil {
				return err
			}
			session.Println(
"[38;2;222;31;171m [38;2;210;31;171m [38;2;207;31;171m [38;2;204;31;171m [38;2;201;31;171m [38;2;198;31;171m [38;2;195;31;171m [38;2;192;31;171m [38;2;189;31;171m [38;2;186;31;171m [38;2;183;31;171m [38;2;180;31;171m [38;2;177;31;171m [38;2;174;31;171m [38;2;171;31;171m [38;2;168;31;171m [38;2;165;31;171m [38;2;162;31;171m [38;2;159;31;171m [38;2;156;31;171m [38;2;153;31;171m [38;2;150;31;171m [38;2;147;31;171m [38;2;144;31;171m [38;2;141;31;171m [38;2;138;31;171m [38;2;135;31;171m [38;2;132;31;171m [38;2;129;31;171m [38;2;126;31;171m [38;2;123;31;171m [38;2;120;31;171m [38;2;117;31;171m [38;2;114;31;171m [38;2;111;31;171m╔[38;2;108;31;171m╦[38;2;105;31;171m╗[38;2;102;31;171m┌[38;2;99;31;171m─[38;2;96;31;171m┐[38;2;93;31;171m╦[38;2;90;31;171m╔[38;2;87;31;171m═[38;2;84;31;171m┬[38;2;81;31;171m [38;2;78;31;171m┬[38;2;75;31;171m╔[38;2;72;31;171m═[38;2;69;31;171m╗[38;2;66;31;171m\r\n",
"[38;2;222;31;171m [38;2;219;31;171m [38;2;216;31;171m [38;2;213;31;171m [38;2;210;31;171m [38;2;207;31;171m [38;2;204;31;171m [38;2;201;31;171m [38;2;198;31;171m [38;2;195;31;171m [38;2;192;31;171m [38;2;189;31;171m [38;2;186;31;171m [38;2;183;31;171m [38;2;180;31;171m [38;2;177;31;171m [38;2;174;31;171m [38;2;171;31;171m [38;2;168;31;171m [38;2;165;31;171m [38;2;162;31;171m [38;2;159;31;171m [38;2;156;31;171m [38;2;153;31;171m [38;2;150;31;171m [38;2;147;31;171m [38;2;144;31;171m [38;2;141;31;171m [38;2;138;31;171m [38;2;135;31;171m [38;2;132;31;171m [38;2;129;31;171m [38;2;126;31;171m [38;2;123;31;171m [38;2;120;31;171m [38;2;117;31;171m [38;2;114;31;171m [38;2;111;31;171m [38;2;108;31;171m║[38;2;105;31;171m [38;2;102;31;171m│[38;2;99;31;171m [38;2;96;31;171m│[38;2;93;31;171m╠[38;2;90;31;171m╩[38;2;87;31;171m╗[38;2;84;31;171m└[38;2;81;31;171m┬[38;2;78;31;171m┘[38;2;75;31;171m║[38;2;72;31;171m [38;2;69;31;171m║[38;2;66;31;171m\r\n",
"[38;2;222;31;171m [38;2;219;31;171m [38;2;216;31;171m [38;2;213;31;171m [38;2;210;31;171m [38;2;207;31;171m [38;2;204;31;171m [38;2;201;31;171m [38;2;198;31;171m [38;2;195;31;171m [38;2;192;31;171m [38;2;189;31;171m [38;2;186;31;171m [38;2;183;31;171m [38;2;180;31;171m [38;2;177;31;171m [38;2;174;31;171m [38;2;171;31;171m [38;2;168;31;171m [38;2;165;31;171m [38;2;162;31;171m [38;2;159;31;171m [38;2;156;31;171m [38;2;153;31;171m [38;2;150;31;171m [38;2;147;31;171m [38;2;144;31;171m [38;2;141;31;171m [38;2;138;31;171m [38;2;135;31;171m [38;2;132;31;171m [38;2;129;31;171m [38;2;126;31;171m [38;2;123;31;171m [38;2;120;31;171m [38;2;117;31;171m [38;2;114;31;171m [38;2;111;31;171m [38;2;108;31;171m╩[38;2;105;31;171m [38;2;102;31;171m└[38;2;99;31;171m─[38;2;96;31;171m┘[38;2;93;31;171m╩[38;2;90;31;171m [38;2;87;31;171m╩[38;2;84;31;171m [38;2;81;31;171m┴[38;2;78;31;171m [38;2;75;31;171m╚[38;2;72;31;171m═[38;2;69;31;171m╝[38;2;66;31;171m\r\n",
"[38;2;222;31;171m [38;2;218;31;172m [38;2;214;31;173m [38;2;210;31;174m [38;2;206;31;175m [38;2;202;31;176m [38;2;198;31;177m [38;2;194;31;178m [38;2;190;31;179m [38;2;186;31;180m [38;2;182;31;181m [38;2;178;31;182m [38;2;174;31;183m [38;2;170;31;184m [38;2;166;31;185m [38;2;162;31;186m [38;2;158;31;187m [38;2;154;31;188m [38;2;150;31;189m [38;2;146;31;190m [38;2;142;31;191m [38;2;138;31;192m [38;2;134;31;193m [38;2;130;31;194m [38;2;126;31;195m [38;2;122;31;196m [38;2;118;31;197m [38;2;114;31;198m [38;2;110;31;199m [38;2;106;31;200m [38;2;102;31;201m [38;2;98;31;202m [38;2;94;31;203m [38;2;90;31;204m [38;2;86;31;205m [38;2;82;31;206m [38;2;78;31;207m [38;2;74;31;208m [38;2;70;31;209m [38;2;66;31;210m [38;2;62;31;211m𝗧[38;2;58;31;212m𝗼[38;2;54;31;213m𝗸[38;2;50;31;214m𝘆[38;2;46;31;215m𝗼[38;2;42;31;216m [38;2;38;31;217m𝗔[38;2;34;31;218m𝗣[38;2;30;31;219m𝗜[38;2;26;31;220m\r\n",
"[38;2;222;31;171m [38;2;219;31;171m [38;2;216;31;171m [38;2;213;31;171m [38;2;210;31;171m [38;2;207;31;171m [38;2;204;31;171m [38;2;201;31;171m [38;2;198;31;171m [38;2;195;31;171m [38;2;192;31;171m [38;2;189;31;171m [38;2;186;31;171m [38;2;183;31;171m [38;2;180;31;171m [38;2;177;31;171m [38;2;174;31;171m [38;2;171;31;171m [38;2;168;31;171m [38;2;165;31;171m [38;2;162;31;171m [38;2;159;31;171m [38;2;156;31;171m [38;2;153;31;171m [38;2;150;31;171m [38;2;147;31;171m [38;2;144;31;171m [38;2;141;31;171m [38;2;138;31;171m╔[38;2;135;31;171m═[38;2;132;31;171m═[38;2;129;31;171m═[38;2;126;31;171m═[38;2;123;31;171m═[38;2;120;31;171m═[38;2;117;31;171m═[38;2;114;31;171m═[38;2;111;31;171m═[38;2;108;31;171m═[38;2;105;31;171m═[38;2;102;31;171m═[38;2;99;31;171m═[38;2;96;31;171m═[38;2;93;31;171m═[38;2;90;31;171m═[38;2;87;31;171m═[38;2;84;31;171m═[38;2;81;31;171m═[38;2;78;31;171m═[38;2;75;31;171m═[38;2;72;31;171m═[38;2;69;31;171m═[38;2;66;31;171m═[38;2;63;31;171m═[38;2;60;31;171m═[38;2;57;31;171m═[38;2;54;31;171m═[38;2;51;31;171m═[38;2;48;31;171m═[38;2;45;31;171m═[38;2;42;31;171m═[38;2;39;31;171m═[38;2;36;31;171m═[38;2;33;31;171m╗[38;2;30;31;171m [38;2;27;31;171m\r\n",
"[38;2;222;31;171m [38;2;219;31;171m [38;2;216;31;171m [38;2;213;31;171m [38;2;210;31;171m [38;2;207;31;171m [38;2;204;31;171m [38;2;201;31;171m [38;2;198;31;171m [38;2;195;31;171m [38;2;192;31;171m [38;2;189;31;171m [38;2;186;31;171m [38;2;183;31;171m [38;2;180;31;171m [38;2;177;31;171m [38;2;174;31;171m [38;2;171;31;171m [38;2;168;31;171m [38;2;165;31;171m [38;2;162;31;171m [38;2;159;31;171m [38;2;156;31;171m [38;2;153;31;171m [38;2;150;31;171m [38;2;147;31;171m [38;2;144;31;171m [38;2;141;31;171m [38;2;138;31;171m║[38;2;135;31;171mH[38;2;132;31;171mT[38;2;129;31;171mT[38;2;126;31;171mP[38;2;123;31;171m-[38;2;120;31;171mT[38;2;117;31;171mO[38;2;114;31;171mK[38;2;111;31;171mY[38;2;108;31;171mO[38;2;105;31;171m [38;2;102;31;171m|[38;2;99;31;171mD[38;2;96;31;171mS[38;2;93;31;171mT[38;2;90;31;171mA[38;2;87;31;171mT[38;2;84;31;171m-[38;2;81;31;171mP[38;2;78;31;171mA[38;2;75;31;171mD[38;2;72;31;171m|[38;2;69;31;171m [38;2;66;31;171mA[38;2;63;31;171mU[38;2;60;31;171mT[38;2;57;31;171mO[38;2;54;31;171mB[38;2;51;31;171mY[38;2;48;31;171mP[38;2;45;31;171mA[38;2;42;31;171mS[38;2;39;31;171mS[38;2;36;31;171m [38;2;33;31;171m║[38;2;30;31;171m\r\n",
"[38;2;222;31;171m [38;2;219;31;171m [38;2;216;31;171m [38;2;213;31;171m [38;2;210;31;171m [38;2;207;31;171m [38;2;204;31;171m [38;2;201;31;171m [38;2;198;31;171m [38;2;195;31;171m [38;2;192;31;171m [38;2;189;31;171m [38;2;186;31;171m [38;2;183;31;171m [38;2;180;31;171m [38;2;177;31;171m [38;2;174;31;171m [38;2;171;31;171m [38;2;168;31;171m [38;2;165;31;171m [38;2;162;31;171m [38;2;159;31;171m [38;2;156;31;171m [38;2;153;31;171m [38;2;150;31;171m [38;2;147;31;171m [38;2;144;31;171m [38;2;141;31;171m [38;2;138;31;171m║[38;2;135;31;171mH[38;2;132;31;171mO[38;2;129;31;171mM[38;2;126;31;171mE[38;2;123;31;171m-[38;2;120;31;171mR[38;2;117;31;171mA[38;2;114;31;171mP[38;2;111;31;171mE[38;2;108;31;171m [38;2;105;31;171m [38;2;102;31;171m|[38;2;99;31;171mU[38;2;96;31;171mD[38;2;93;31;171mP[38;2;90;31;171mR[38;2;87;31;171mA[38;2;84;31;171mW[38;2;81;31;171m [38;2;78;31;171m [38;2;75;31;171m [38;2;72;31;171m|[38;2;69;31;171m [38;2;66;31;171mC[38;2;63;31;171mL[38;2;60;31;171mD[38;2;57;31;171mA[38;2;54;31;171mP[38;2;51;31;171m [38;2;48;31;171m [38;2;45;31;171m [38;2;42;31;171m [38;2;39;31;171m [38;2;36;31;171m [38;2;33;31;171m║[38;2;30;31;171m [38;2;27;31;171m [38;2;24;31;171m\r\n",
"[38;2;222;31;171m [38;2;219;31;171m [38;2;216;31;171m [38;2;213;31;171m [38;2;210;31;171m [38;2;207;31;171m [38;2;204;31;171m [38;2;201;31;171m [38;2;198;31;171m [38;2;195;31;171m [38;2;192;31;171m [38;2;189;31;171m [38;2;186;31;171m [38;2;183;31;171m [38;2;180;31;171m [38;2;177;31;171m [38;2;174;31;171m [38;2;171;31;171m [38;2;168;31;171m [38;2;165;31;171m [38;2;162;31;171m [38;2;159;31;171m [38;2;156;31;171m [38;2;153;31;171m [38;2;150;31;171m [38;2;147;31;171m [38;2;144;31;171m [38;2;141;31;171m [38;2;138;31;171m║[38;2;135;31;171mL[38;2;132;31;171mD[38;2;129;31;171mA[38;2;126;31;171mP[38;2;123;31;171m [38;2;120;31;171m [38;2;117;31;171m [38;2;114;31;171m [38;2;111;31;171m [38;2;108;31;171m [38;2;105;31;171m [38;2;102;31;171m|[38;2;99;31;171mA[38;2;96;31;171mC[38;2;93;31;171mK[38;2;90;31;171mR[38;2;87;31;171mA[38;2;84;31;171mW[38;2;81;31;171m [38;2;78;31;171m [38;2;75;31;171m [38;2;72;31;171m|[38;2;69;31;171m [38;2;66;31;171mO[38;2;63;31;171mV[38;2;60;31;171mH[38;2;57;31;171m-[38;2;54;31;171mT[38;2;51;31;171mO[38;2;48;31;171mK[38;2;45;31;171mY[38;2;42;31;171mO[38;2;39;31;171m [38;2;36;31;171m [38;2;33;31;171m║[38;2;30;31;171m\r\n",
"[38;2;222;31;171m [38;2;219;31;171m [38;2;216;31;171m [38;2;213;31;171m [38;2;210;31;171m [38;2;207;31;171m [38;2;204;31;171m [38;2;201;31;171m [38;2;198;31;171m [38;2;195;31;171m [38;2;192;31;171m [38;2;189;31;171m [38;2;186;31;171m [38;2;183;31;171m [38;2;180;31;171m [38;2;177;31;171m [38;2;174;31;171m [38;2;171;31;171m [38;2;168;31;171m [38;2;165;31;171m [38;2;162;31;171m [38;2;159;31;171m [38;2;156;31;171m [38;2;153;31;171m [38;2;150;31;171m [38;2;147;31;171m [38;2;144;31;171m [38;2;141;31;171m [38;2;138;31;171m║[38;2;135;31;171mO[38;2;132;31;171mV[38;2;129;31;171mH[38;2;126;31;171m-[38;2;123;31;171mB[38;2;120;31;171mE[38;2;117;31;171mA[38;2;114;31;171mM[38;2;111;31;171m [38;2;108;31;171m [38;2;105;31;171m [38;2;102;31;171m|[38;2;99;31;171mN[38;2;96;31;171mF[38;2;93;31;171mO[38;2;90;31;171m-[38;2;87;31;171mT[38;2;84;31;171mO[38;2;81;31;171mK[38;2;78;31;171mY[38;2;75;31;171mO[38;2;72;31;171m|[38;2;69;31;171m [38;2;66;31;171mN[38;2;63;31;171mF[38;2;60;31;171mO[38;2;57;31;171m-[38;2;54;31;171mC[38;2;51;31;171mL[38;2;48;31;171mA[38;2;45;31;171mP[38;2;42;31;171m [38;2;39;31;171m [38;2;36;31;171m [38;2;33;31;171m║[38;2;30;31;171m\r\n",
"[38;2;222;31;171m [38;2;219;31;171m [38;2;216;31;171m [38;2;213;31;171m [38;2;210;31;171m [38;2;207;31;171m [38;2;204;31;171m [38;2;201;31;171m [38;2;198;31;171m [38;2;195;31;171m [38;2;192;31;171m [38;2;189;31;171m [38;2;186;31;171m [38;2;183;31;171m [38;2;180;31;171m [38;2;177;31;171m [38;2;174;31;171m [38;2;171;31;171m [38;2;168;31;171m [38;2;165;31;171m [38;2;162;31;171m [38;2;159;31;171m [38;2;156;31;171m [38;2;153;31;171m [38;2;150;31;171m [38;2;147;31;171m [38;2;144;31;171m [38;2;141;31;171m [38;2;138;31;171m║[38;2;135;31;171m1[38;2;132;31;171m0[38;2;129;31;171m0[38;2;126;31;171mU[38;2;123;31;171mP[38;2;120;31;171m-[38;2;117;31;171mS[38;2;114;31;171mL[38;2;111;31;171mA[38;2;108;31;171mP[38;2;105;31;171m [38;2;102;31;171m|[38;2;99;31;171mF[38;2;96;31;171mN[38;2;93;31;171m-[38;2;90;31;171mT[38;2;87;31;171mO[38;2;84;31;171mK[38;2;81;31;171mY[38;2;78;31;171mO[38;2;75;31;171m [38;2;72;31;171m|[38;2;69;31;171m [38;2;66;31;171mA[38;2;63;31;171mP[38;2;60;31;171mE[38;2;57;31;171mX[38;2;54;31;171m-[38;2;51;31;171mB[38;2;48;31;171mE[38;2;45;31;171mA[38;2;42;31;171mM[38;2;39;31;171m [38;2;36;31;171m [38;2;33;31;171m║[38;2;30;31;171m\r\n",
"[38;2;222;31;171m [38;2;219;31;171m [38;2;216;31;171m [38;2;213;31;171m [38;2;210;31;171m [38;2;207;31;171m [38;2;204;31;171m [38;2;201;31;171m [38;2;198;31;171m [38;2;195;31;171m [38;2;192;31;171m [38;2;189;31;171m [38;2;186;31;171m [38;2;183;31;171m [38;2;180;31;171m [38;2;177;31;171m [38;2;174;31;171m [38;2;171;31;171m [38;2;168;31;171m [38;2;165;31;171m [38;2;162;31;171m [38;2;159;31;171m [38;2;156;31;171m [38;2;153;31;171m [38;2;150;31;171m [38;2;147;31;171m [38;2;144;31;171m [38;2;141;31;171m [38;2;138;31;171m║[38;2;135;31;171mR[38;2;132;31;171m6[38;2;129;31;171m-[38;2;126;31;171mG[38;2;123;31;171mR[38;2;120;31;171mE[38;2;117;31;171mE[38;2;114;31;171mN[38;2;111;31;171m [38;2;108;31;171m [38;2;105;31;171m [38;2;102;31;171m|[38;2;99;31;171mR[38;2;96;31;171mO[38;2;93;31;171mB[38;2;90;31;171mL[38;2;87;31;171mO[38;2;84;31;171mX[38;2;81;31;171m [38;2;78;31;171m [38;2;75;31;171m [38;2;72;31;171m|[38;2;69;31;171m [38;2;66;31;171mC[38;2;63;31;171mO[38;2;60;31;171mD[38;2;57;31;171m-[38;2;54;31;171mA[38;2;51;31;171mL[38;2;48;31;171mL[38;2;45;31;171m [38;2;42;31;171m [38;2;39;31;171m [38;2;36;31;171m [38;2;33;31;171m║[38;2;30;31;171m [38;2;27;31;171m\r\n",
"[38;2;222;31;171m [38;2;219;31;171m [38;2;216;31;171m [38;2;213;31;171m [38;2;210;31;171m [38;2;207;31;171m [38;2;204;31;171m [38;2;201;31;171m [38;2;198;31;171m [38;2;195;31;171m [38;2;192;31;171m [38;2;189;31;171m [38;2;186;31;171m [38;2;183;31;171m [38;2;180;31;171m [38;2;177;31;171m [38;2;174;31;171m [38;2;171;31;171m [38;2;168;31;171m [38;2;165;31;171m [38;2;162;31;171m [38;2;159;31;171m [38;2;156;31;171m [38;2;153;31;171m [38;2;150;31;171m [38;2;147;31;171m [38;2;144;31;171m [38;2;141;31;171m [38;2;138;31;171m║[38;2;135;31;171m [38;2;132;31;171m [38;2;129;31;171m [38;2;126;31;171m [38;2;123;31;171m [38;2;120;31;171m [38;2;117;31;171m [38;2;114;31;171m [38;2;111;31;171m [38;2;108;31;171m [38;2;105;31;171m [38;2;102;31;171m|[38;2;99;31;171m2[38;2;96;31;171mK[38;2;93;31;171m-[38;2;90;31;171mB[38;2;87;31;171mE[38;2;84;31;171mA[38;2;81;31;171mM[38;2;78;31;171m [38;2;75;31;171m [38;2;72;31;171m|[38;2;69;31;171m [38;2;66;31;171m [38;2;63;31;171m [38;2;60;31;171m [38;2;57;31;171m [38;2;54;31;171m [38;2;51;31;171m [38;2;48;31;171m [38;2;45;31;171m [38;2;42;31;171m [38;2;39;31;171m [38;2;36;31;171m [38;2;33;31;171m║[38;2;30;31;171m [38;2;27;31;171m [38;2;24;31;171m [38;2;21;31;171m\r\n",
"[38;2;222;31;171m [38;2;219;31;171m [38;2;216;31;171m [38;2;213;31;171m [38;2;210;31;171m [38;2;207;31;171m [38;2;204;31;171m [38;2;201;31;171m [38;2;198;31;171m [38;2;195;31;171m [38;2;192;31;171m [38;2;189;31;171m [38;2;186;31;171m [38;2;183;31;171m [38;2;180;31;171m [38;2;177;31;171m [38;2;174;31;171m [38;2;171;31;171m [38;2;168;31;171m [38;2;165;31;171m [38;2;162;31;171m [38;2;159;31;171m [38;2;156;31;171m [38;2;153;31;171m [38;2;150;31;171m [38;2;147;31;171m [38;2;144;31;171m [38;2;141;31;171m [38;2;138;31;171m╚[38;2;135;31;171m═[38;2;132;31;171m═[38;2;129;31;171m═[38;2;126;31;171m═[38;2;123;31;171m═[38;2;120;31;171m═[38;2;117;31;171m═[38;2;114;31;171m═[38;2;111;31;171m═[38;2;108;31;171m═[38;2;105;31;171m═[38;2;102;31;171m═[38;2;99;31;171m═[38;2;96;31;171m═[38;2;93;31;171m═[38;2;90;31;171m═[38;2;87;31;171m═[38;2;84;31;171m═[38;2;81;31;171m═[38;2;78;31;171m═[38;2;75;31;171m═[38;2;72;31;171m═[38;2;69;31;171m═[38;2;66;31;171m═[38;2;63;31;171m═[38;2;60;31;171m═[38;2;57;31;171m═[38;2;54;31;171m═[38;2;51;31;171m═[38;2;48;31;171m═[38;2;45;31;171m═[38;2;42;31;171m═[38;2;39;31;171m═[38;2;36;31;171m═[38;2;33;31;171m╝[38;2;30;31;171m [0;00m\r\n",
"ALL METHODS FOR { TOKYO API } ARE IN \"UPPER CASE\"\r\n",
)
			return nil
		},
	})
}