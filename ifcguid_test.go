package ifcguid

import (
	"bufio"
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_ToUuid_and_FromUuid(t *testing.T) {
	testData := `
01cf62c8-e9bc-bf88-0000-000000000005;01psB8wRo$Y00000000005
3085A8E4-61FD-4776-9FF1-1B24A646CA4F;0mXQZaOVr7Tf$n6oIcHifF
12197C0B-DFA7-4C19-B3E6-D1A9A59663AB;0I6NmBtwTC6RFcqQcbbcEh
163EFED7-1B9A-4E8C-B69A-6497F95C232A;0MFlxN6vfEZBQQP9VvN2Cg
182CA57E-0C83-48CC-A47A-8514F066DEB2;0OBAL_38D8pAHwXHJmPjwo
9BB6C3B7-A412-41A9-9851-4BCCB0ED0526;2RjiEtf191gPXHIyomxGKc
5B2350DF-CE5E-482F-9ACF-0BCA06C4E4D7;1R8r3Vpbv8BvhF2ye6nEJN
E862FFB0-6793-4340-AD5E-62D18A70F9E3;3eOl_mPvD3GArUOj6ASFdZ
C88129C7-3F45-46E8-AFFF-EAC60FEDE0BE;38WId7FqL6wA$$wiOFxU2_
4C155151-2642-4FA1-AB0A-80D2BAF5FA89;1C5L5H9a9FeQiAWDAwzVg9
B84A4E66-97FB-4FA2-B1FE-2B230171170C;2uIavcb$jFeh7_AoC1SHSC
8F64282D-4B0B-4A75-B798-254A54A59DEF;2FP2WjImjATRUO9KfKfPtl
66BDA932-AC14-40B8-BA02-630349561ABF;1clQaoh1H0kBe2OmD9LXg$
59EB0A89-2B10-4A22-A642-EEC4049B831E;1Pwmg9An1A8gP2xiG4cuCU
348771A0-2E84-4C24-96D4-6780F633D002;0qXt6WBeHC99RKPu3sCz02
A9B50BD5-0EBA-4072-AC63-FC7484577A3E;2fjGlL3hf0SgnZ$7I4Lte_
D3D61E31-4CD8-4C08-8C0A-9A278AAA9B4E;3JrXunJDXC28mAcYUAgfjE
4E115A05-2D0B-4976-B811-F724FC399C41;1E4Le5BGj9ThWHzoJyEPn1
6616483D-07E3-4E27-90B5-A54C50440E32;1c5aWz1_DE9v2rfKnGH0uo
0E6A8F85-27C8-4BF1-9F07-87118918D480;0EQe_59yXByPy7Xn696DI0
00C279B5-E1E6-476E-B087-A9777B77253E;00mdcruUP7Rh27gNTxToK_
B74FE921-BB2E-44F9-8A83-F5AAAC62108D;2tJ_aXkov4_Og3zQgiOX2D
1C6CBEF4-2F51-4058-B02A-B5AA22DFEA42;0SRBxqBr50MB0gjQeYt_f2
2FFE4C67-12C2-488E-82DB-996638F29770;0l$and4i98ZeBRcMOuyfTm
233E7E68-1E33-4AA0-97DD-E6D485CF8E6E;0ZFdve7ZDAe9VTvjI5puvk
114ED2EB-CB57-43D3-B252-CEE041FCDC1F;0HJjBhorT3qx9Ipk11$DmV
BE9DDB66-B4DC-45CE-8BCB-FEA9A3A71E69;2_dTjcjDn5pelB$gcZfnvf
88778FDD-F823-4B90-A68A-B1C14BD9E356;28Tu$T_2DBaAQAiS5BsUDM
AD207452-482C-44CA-9070-8ADE649A16ED;2j87HII2n4of1mYjvacXRj
4D8F7313-D06E-45E4-8182-5E2F4F81FE48;1DZtCJq6v5v862NYzFWVv8
01973BAE-7066-4702-BFC9-486D7637201C;01bpkkS6P70h$9I6rsDo0S
8EA2CDD2-6662-41CE-863C-6AD8840C03DB;2EeitIPc91peOyQjY430FR
3FAF53B7-E9AB-4D5C-916A-2344C7E829AA;0$hrEtwQjDN95g8qJ7w2cg
B4CECE78-400D-4C78-AF3D-3B4865BE9D2E;2qpivuG0rCUAyzEqXblfqk
1D0E7B96-27C3-4388-86D1-349B297969F7;0T3dkM9yD3Y8RHD9ifUMdt
13CBD3A9-D450-4120-A6EE-40B0AD3F4BA3;0JozEfr5118ARkGB2jFqkZ
09702D64-0121-49DD-B851-ADA9A847C222;09S2ra0I59tRXHhQceHy8Y
277DDB9B-0F7C-4CD8-9962-A1D589B386B9;0dVTkR3tnCs9bYeTM9iuQv
9E4FA830-373E-4B4A-AFC8-92ED51CFC007;2UJwWmDpvBIg$8akrHpy07
4D5F5AC0-E636-425A-B05E-9FDC9F2DDE6A;1DNrh0vZP2Mh1UdzoVBTvg
DDE8E6EA-BDEB-4920-8601-F2BDEB402B0D;3TwERglUj988O1yhthG2iD
75145D75-AF9D-4BF4-BE33-25686C148FF5;1r55rrhvrBzBup9MXi58$r
7243DF56-F446-4BAC-8D1F-FAFD8F5DE9BA;1oGzzMz4PBh8qV_lsFNUcw
4A83EE5D-6C43-42F0-9E3E-4704A0BEE2E1;1AW_vTR4D2y9u_HmIWlkBX
90ABDB53-553B-4F0A-B977-196F767CCD2B;2GgzjJLJjF2hbt6MzsVCqh
F93A9570-7AD6-465E-B00B-E76451F5B4F6;3vEfLmUjP6Nh0BvsHHzRJs
13519721-935D-4DFA-B5C4-700185639895;0JKPSXarrD_hN4S065OvYL
B91206A2-D5E6-47C3-BBC6-333D02899BD2;2v4WQYrUP7mxl6Cpq2YPlI
1E0D8434-5C90-4738-921E-048791073B7B;0U3OGqN917E98U18UH1pjx
94584F5D-47BE-4B50-9700-4744407CF46D;2KM4zTHxvBK9S0HqH0VFHj
E193E5BD-83ED-4C63-92A8-0180FCB53C7F;3Xa_MzW_rCOvAe0O3yjJn$
E8A22D9E-DB58-4B60-B122-6A9D202F1852;3eeYsUsrXBOB4YQfqWBnXI
BD47B7CE-3DF0-450E-878C-AED5C46000F3;2zHxVEFV153eUChjN4O03p
C038590D-CB2A-41A3-88F5-CAC3A30147DE;30E5aDoof1euZroiEZ0KVU
AD574B6C-9EFA-4B64-8F39-E7D54C36AF4D;2jLqjidlfBP8yvvzLCDgzD
902C52A9-EC73-4C60-A96C-6E256E9FE662;2GB5Afx7DCOAbiRYLkd_PY
2F3A801D-138E-479D-948E-BFF44085FBA4;0lEe0T4uv7dPIEl$H0XVka
92960212-3D88-42C8-9518-37B4215B5807;2IbW8IFOX2o9KODxGXMrW7
1BE6F6BF-1270-44E4-B35C-56007CB9C3E7;0RvlQ$4d14vBDSLW1ykSFd
BA45794C-C44E-4059-AC31-938121C6AA01;2wHNbCn4v0MQmnau4Xnge1
52D531C6-400F-4BB4-9FB7-0EC3A331F5CB;1IrJ76G0zBj9_t3iEZCVNB
2F5AD2D1-D3ED-4F08-971D-9E3F8C949CE0;0lMjBHq_rF29STdZ_Cb9pW
DD0F41B0-8604-4662-9DFB-95F8B9026716;3T3q6mXWH6OftxbVYv0cSM
688D849F-9755-4F5D-B6D9-787AC4695F1C;1eZOIVbrLFNRRPU7h4QLyS
25643E0E-3AED-484D-B302-8B1C75CA7ADC;0bP3uEEkr8JRC2YnnrodhS
2E1B9BDE-12E5-427D-92CE-21E87213DB8A;0k6vlU4kL2VPBE8UXo4zkA
536D9E3D-EE5B-43C6-9BF2-E2F965151C5F;1JRPuzxbj3nfloulbb5HnV
34CBA891-F557-41A4-A172-BD8558C13ADC;0qowYHzLT1fA5olOLOmJhS
46B92BA0-01FD-4309-88E3-F2364DB3B024;16kIkW0Vr32OZZyZPDix0a
BDB2AC77-E807-4F0B-B031-D8E8997BA3FC;2zigntw0TF2x0nsEYPUwFy
794BDA5D-0AF7-45B1-9AEF-171004AB4630;1vIzfT2lT5iPhl5n04gqOm
D28EA8B0-9026-4534-934F-C7D1786E56F5;3IZgYma2P5D9DFnz5uRbRr
4ADBF3B4-4BD0-46E8-B398-C87E1567BAE4;1As$EqIz16wBEOo7uLPxha
6CAFF339-4EB0-410B-9D55-233B541AEF14;1ih$CvJh112vrL8pjK6kyK
F8D3E203-3807-4BAD-A110-3300692A9782;3uq_83E0TBhQ4GCm1fAfU2
8E736F9D-243B-4EDC-A9A7-226C2B6585FD;2ESs_T93jEtAcd8cmhPONz
4039E7F9-4C49-4B9D-B1C5-1A6B5256E578;10EUVvJ4bBdR756cjILkLu
8A57C7EC-16DF-4B77-B3FD-6CE284AFC0C8;2ALyVi5jzBTxFzREA4hy38
933C3E03-B196-463F-892E-316A293AF0DE;2JF3u3iPP6FuakCMefEl3U
08B3FABE-D4A3-4E1C-A2F1-2A0B41B2240C;08i$g_rADE7ABnAWj1iYGC
26E330B9-FFAE-48C5-AE7E-DBAC6D514C05;0cup2v$wv8nQv_swnjKKm5
72D39AA1-ABF0-46DE-ADC4-F941707CB3A3;1oqvgXg$16tgt4_K5mVBEZ
8383D0A9-A246-47CF-B228-F777AE26E793;23Wz2feaP7px8eztUk9kUJ
0D432EC4-BECE-4008-93AF-5FAC9D4D8954;0DGox4liv029ElNwoTJObK
5B78B068-0F9E-4372-A8CF-F5E56673BFC3;1RUB1e3vv3SgZFzULcSx$3
411BD9F8-6D2F-4423-8851-10DBBD34E862;116zduRIz48uXH4DkzDEXY
238C27DE-79F0-436F-A0CD-503BDB1B8C70;0ZZ2VUUV13Rw3DK3lR6unm
13B109F8-3BB8-4A67-BDA5-88169ED8F71E;0JiGduExXAPxsbY1QUsFSU
AA4393DD-A1AD-43E9-B703-71D763166FF8;2gGvFTeQr3wRS3STTZ5c$u
27E6A1FA-72E9-4C79-AFA9-9D3DC748BB33;0dvg7wSkbCUQ_fdJt7IBip
25C689FD-2C12-4203-B3E0-244616A4C4BE;0bnedzB1920xFW94OMfCI_
B574A45B-5870-44B1-8405-3DB62FEEA488;2rTAHRM714iOG5FROlxgI8
6E7AA299-9FBC-4716-ACDA-32F64A79F8DD;1kUgAPdxn75gpQClPAUVZT
B68BED65-91C1-48F3-9B4F-019DBBCC4390;2sY_rbaS58yvjF0Psxp4EG
FD811263-2B8C-4DA1-991D-A58FD9FA6791;3zWH9ZAunDePaTfO$P_cUH
AB8C6E0D-0E70-449A-8618-76324E8549CE;2hZ6uD3d14ceOOTZ9EXKdE
436B2BF4-17FC-417A-A58B-9173BE02CA72;13Qolq5$n1UgMBaNE_0ifo
ECC05584-4989-4E64-9A65-5829AC478397;3im5M4IObEP9fbM2ciHuEN
D783F488-1E65-4A34-80B3-956A6E9090E5;3NW$I87cLAD82pbMfka93b
EFEE0A6F-2638-4694-804F-D2F08D1E7F7B;3lxWfl9ZX6b81Fql2D7dzx
11AD6281-B704-44EB-82EE-59337F5BDBA5;0HhMA1jmH4wuBkMJD$Mzkb
08A20574-612F-43E7-B675-0C9C4C64446C;08eWLqOIz3vxPr39nCP4Hi
FAFE2D42-3D32-4411-B5B3-E9B467CDBF0B;3w$Yr2FJ944RMpwRHdpRyB
E8EB1BFC-A398-4B19-8DFA-6F53B62FA11F;3ewnlyevXB6OtwRrEsBw4V
04B3B6FE-14C4-462C-B923-E6D336565F87;04ixR_5CH6BBaZvjCsLb_7
BD372319-877E-408F-A97D-0B63ECA20478;2zDoCPXtv0Zwbz2sFieWHu
27B93FD7-D506-4FC5-B306-1A3DB57B9363;0dkJ$NrGPFnRC66ZsrUvDZ
167F99F6-7C6C-4C92-B99E-2A0A372D2FA0;0MVvdsV6nCahcUAWetBI_W
04E3AC5F-710D-43B1-A989-9020FC84F331;04uwnVSGr3iQc9a23yXFCn
647F1044-E6AB-4D3C-A952-F98FC7E23502;1aVn14vgjDFAbI_O$7uZK2
748774AF-6751-411D-AD06-F51D7E85965F;1qXtIlPr517Qq6zHr_XPPV
7DAFE0FA-6768-4269-B99B-1649873FD3E8;1zh_3wPsX2QRcR5ac7FzFe
03DFFB93-C814-4CB6-AFC0-9C7FC71A11D5;03t$kJo1HCjg$0d7$76X7L
CB9456F9-1826-4C63-A81D-6B0F30232DFC;3Bb5Rv62PCOwWTQmym8oty
F62F35A0-8CE3-48DF-A9BA-1D771CE44E66;3sBpMWZED8twcw7NSSv4vc
3C5D0C4D-C359-4231-BAB4-B425994DC0C8;0yNGnDmrb2CRgqj2MPJS38
42B64AFF-61CD-4F41-83D2-1B963FF66834;12jah$OSrFGOFI6vO$zcWq
400A4619-C518-4EB7-91A4-93D012D49A1B;102aOPnHXEjv6aaz0Ir9eR
E4C45EEC-8C7C-483B-9FFD-1D5A5A25AF5B;3an5xiZ7n8Ev$z7LfQ9QzR
2FE9508B-F512-49B7-B750-D97EA8F2B8A4;0lwL2BzH99jxTGsNweyhYa
57088090-F6A0-4100-A666-9AB72178E1B2;1N282Gzg110APcchSXUE6o
529266FF-4F13-41A5-A050-185A936EEF80;1IacR$JnD1fQ1G65gJRk_0
FBD923A9-465F-4779-AE53-BD813547B00E;3xsIEfHbz7UQvJlO4rHx0E
CEA96A19-3BC4-4869-959C-A897D7618118;3EgMePEyH8QPMSg9VNOO4O
C0EA9E99-E44F-4FBC-9C83-F040804687E5;30wfwPv4zFl9o3y420HeVb
91E0F572-C6D7-4A05-8582-A1FFB2ED617C;2HuFLonjTA1OM2eV_oxM5y
FEB5765F-C8E4-44A4-96C5-70101260B441;3_jNPVoEH4f9R5S10IOBH1
73850CCD-6FC6-4A64-A844-A2BC6CF58E8A;1pXGpDRyPAPAX4ehnizOwA
ED60B83B-BADA-464B-A50D-18807EFD8A9E;3jOBWxkjf6IwKD681_$OgU
76E963F8-7A60-4CAA-A4C7-3433BC7415A3;1swMFuUc1CggJ7D3EyT1MZ
9CE8366E-6318-4E0D-847B-5352225B0A28;2Sw3PkOnXE3OHxKr8YMmee
2410B058-EC42-4E79-A526-E4FE4267D391;0a4B1Ox49EUQKcvFv2PzEH
BDE89939-EB73-48A3-A85F-76A170F870D0;2zw9avwtD8ewXVTg5m_73G
E5C6C90A-025F-4667-A2F9-AA082A0A81C6;3bniaA0bz6PwBvgWWg2e76
CF84935B-C0B3-416B-8285-E6009318BBF6;3FX9DRmBD1QuA5vW2J6Bls
71C11D4B-EE24-4D6F-AFDE-A7F86D8C9DF2;1nmHrBxYHDRw$Uf$XjZ9to
CF224B37-B0A2-4556-A6EF-04A06E6D22FD;3F8aitiA95LgRl1A1kRIBz
7489BD9E-BDDF-47EF-9E29-A864E38BE43A;1qYRsUlTz7xvufg6JZY_Gw
6624AE3E-5680-485E-AA98-FB3845CDE132;1c9Au_Le18NggO_pX5pU4o
C2112D4D-E307-45C7-B42A-04E94795C824;324IrDumT5nxGg1Eb7bSWa
25A1BD60-EFBB-4A23-995E-5A8E578CD3E6;0beRrWxxjA8vbUMevNZDFc
47C94809-84F2-47D3-8550-3D9A934BFB5B;17oKW9XF97quLGFPgJI$jR
61869E89-274D-4635-A499-EF1E7F9EF3D7;1XXfw99qr6DQIPxnv$dlFN
C717E0C8-91FB-43EA-B083-DDDDAF74AFD6;375_38aVj3wh23tTslTA$M
BC7E5E5C-6DA8-4D1B-9396-B554D5F79134;2yVbvSRQXD6vEMjLJLzv4q
596FC212-0654-47D9-9719-DE402EB1B290;1PRy8I1bH7sPSPta0kiRAG
FEFC1E52-57EF-4BD9-8187-76C534D202C7;3_$1vIL_zBsO67TiKqqWB7
440965CE-8D3F-43F7-BCA6-6C283160A413;142MNEZJz3zxocR2WnOAGJ
02210FFF-0D41-43DB-8170-06F57B26A279;028G$$3K53su5m1lLx9g9v
BF19D723-6C40-4517-9DD6-4C64CACE5653;2$6TSZR4155vtMJ6JApbPJ
355A9131-8615-4789-BEF3-7929ED778CC9;0rMf4nXXL7YRxpUIdjTup9
D157D5F8-0442-4F11-B6B1-D541D6768729;3HLzNu149F4RQnrK7MTeSf
01ABD14B-289B-4929-81D4-F1A23FA2D51B;01gz5BA9j9AO7KyQ8$ejKR
59908B2A-6F80-46D0-8DCB-96C7E3A311D3;1Pa8igRu16q8tBbiVZen7J
E4D34551-0479-44B3-BA45-EEBA9F02AC37;3aqqLH17b4ixf5xhgV0gmt
AA8DA04D-D32A-4BB8-A76F-0E2102016D4F;2gZQ1DqofBkATl3Y420MrF
8F110226-20A5-4F88-A841-B68D12C7606F;2F4G8c8ALFYAX1jeqIns1l
8A4506F8-9277-40DB-9576-B6D9E7D67C76;2AHGRuadT0svLsjjddrdns
B8148D71-A0C1-4666-8D65-8FCBB118845E;2u58rneC56PerbZykn68HU
F1BB68AE-CE2D-4749-AB43-E177A9A6DAD7;3nksYkpYr7IQj3uNUffjhN
B84FCE5D-F4C9-4D03-A05E-1C3EB6E53C02;2uJyvTzCbD0w1U73wsvJm2
EC722D42-742D-4EBF-A649-82D3F5996297;3iSYr2T2rElwP9WjFrcMAN
8DD86B8A-43A9-4291-A343-989A3452FCAA;2Ds6kAGwb2aQD3c9eqKlog
42F97FB6-16A8-4B15-9AF3-51B2ED9C0196;12_N_s5gXB5PhpKRBjd06M
1291AD5B-2007-47E3-A85A-8EADDA067CCA;0IaQrR80T7uwXQZgtQ1dpA
250A7F16-0080-4067-A24A-9E6A707B94A8;0b2dyM0810Pw9AdcfmUvIe
983D81C9-3F2E-4B79-9433-0E72458811FD;2OFO79FovBUPGp3d95Y17z
1D73DBE9-4CBA-450B-B6B7-4F781B48957A;0TSzlfJBf52xQtJtWRI9Lw
DD8ACB75-266D-421C-9F28-03FB65B47E17;3TYijr9cr279ye0$jbj7uN
9F87A3B7-C6A6-4E49-ACFD-619027BC5399;2VXwEtngPEIQpzOP0dl5EP
01F528B8-1360-4F02-84BE-9BCE127B3BDA;01zIYu4s1F0eI_cyuIUplQ
B0079989-A6DA-45CE-BC95-ED4B6DF42EE3;2m1vc9fjf5phoLxKjjz2xZ
219CF4BA-08CA-412E-BF7B-D3FF728314D8;0XdFIw2Cf1Bhzxq$zoWnJO
CAC10F31-4194-4A99-A887-B1BE609FDEBA;3AmGynGPHAcQY7iRvWdzww
81B28414-D19A-41AF-8550-3AE0A8861CA0;21ieGKqPf1huLGEk2eXXoW
FC664F9C-263C-4F0E-9D43-00DA78E5403B;3yPa_S9ZnF3fr30DfuvK0x
A346E04E-DF77-477E-93C9-27C83B6B4706;2ZHk1EttT7VfF99yWxQqS6
607958F2-9582-4320-9A4A-DDB10A0E55CD;1WULZobO9389fAtR4A3bND
2A6352B8-3FBC-476E-96DC-6851883F215A;0gOrAuFxn7RfRSQ568Fo5Q
14462A71-7CB9-41B3-955C-E0B2486FDD92;0KHYfnVBb1ivLSuB98RzsI
406BE8FE-2084-456A-82CF-D930E82615B0;10Q_Z_88H5QeBFsJ3e9XMm
60C013D3-5324-497E-B4C7-264610981235;1Wm1FJKoH9VhJ79aOGc18r
4F379773-EBD6-45CF-A342-227A88B83EA0;1FDvTpwzP5pwD28dg8k3wW
240603EB-A928-40E6-83A5-F5D4FDCA1C7B;0a1WFhgIX0veEbzTJzoXnx
21D4B317-7ECA-48F0-A903-A00F7CDFBD52;0XrBCNVif8yAa3e0zytxrI
3E3C5D6D-22D1-4653-AF7D-C51E92E1A016;0_F5rj8j56KwzznHwIuQ0M
03842665-771F-4C8A-88E4-BE1BDE6117EB;03X2PbTnzCYeZalXlUOHVh
632E586C-BF9C-47D9-94B6-D12DF92DCD54;1ZBbXilvn7sPIsqItvBSrK
86B4F76F-484F-4C82-899F-EE4BAAA7806D;26jFTlI4zCWecVxakgfu1j
EF379F1C-034C-4918-81BF-D3A26378460C;3lDvyS0qn9686$qw9ZU4OC
676B2C07-1914-4734-B3FE-5BCA665E3965;1dQom76HH7DBF_MyfcNZbb
57EE9C3A-22D3-41B6-95EF-29101194FAEE;1Nxfmw8jD1jfNlAH0HbFhk
A4A6727D-A4F2-4380-A31A-17C1D7A8BFCE;2afd9zfF93WACQ5y7NgB$E
FC101646-8CE3-42BD-B41A-46DCCAFC940C;3y41P6ZED2lRGQHjpA$9GC
B760E6C5-FE70-48BF-9C5D-933B0FBD03CC;2tOER5$d18lvnTapiFlGFC
63824448-2CAC-41FF-98C9-E0481161325A;1ZWaH8BAn1$vZ9u4WHOJ9Q
D6561D74-DDA7-4AAE-B0C9-BBAB4A861B42;3MLXrqtQTAhh39kwjAXXj2
93BA81AF-164C-41CE-BE7A-1B53B1A34197;2Jke6l5an1phvw6rEneq6N
F62D9960-7344-4EEE-98C4-43306A0DD63D;3sBPbWSqHExfZ4Gp1g3TOz
41B8BBCA-A459-4635-BD5D-AE2C61CD42D7;11kBlAf5b6DRrThYnXpKBN
6674C8B6-3EDC-4DB5-9620-BA0DBACE18C3;1cTCYsFjnDjPOWkWswpXZ3
5E25CC1D-F4CF-4CC7-85E3-B28855992431;1U9SmTzCzCnuNZieXLcIGn
2DA1095D-E32A-421C-B0AA-5D4934FEBB46;0jeGbTuof27B2gNKaq$hj6
DC7B7FD2-9E36-446A-AD1F-7CDD6E5D58F6;3SUt$IdZP4QgqVVDrkNLZs
E067931B-3733-46B4-96E6-95E2CDF34963;3WPvCRDpD6j9RcbUBDyqbZ
7F8D6777-ACB5-4935-8828-345746B210D4;1$ZMTthBL9DOWeD5T6iX3K
521C6B69-8CE4-4714-A44C-641679D8F801;1I76jfZEH75AHCP1PvsFW1
DF340BCF-F292-40B5-8826-A0D5A64FA664;3VD0lFyf90jOWceDMcJwPa
9E53D3B9-8F02-4A23-9092-9F51B5970F86;2UKzEvZm9A8v2Idr6rbm_6
D53E4BAB-D127-4A52-B8DA-05908EF77645;3LFakhqITAKhZQ1P2EztP5
6F18F6B4-C970-4726-BDDA-02DA0078F7AC;1l6FQqoN179htQ0je0UFUi
E929F69C-234A-475F-9ED8-843A93866401;3fAVQS8qf7NvxOX3gJXcG1
CEF97CFA-EE61-4795-9CDC-F8E0F0F88AB6;3E_Npwxc57bPpS_E3m_8gs
6E062D7A-98C8-4295-BB16-F2D1AA1C162E;1k1YrwcCX2bRiMyj6g71Ok
C3FF8931-8304-479C-8300-CED33D6A32FF;33$uanWmH7d8C0pjCzQZB$
BAA1CC9B-4616-4932-960A-9697A5E9EE99;2weSoRHXP9CfOAbfUbwUwP
C123100E-D157-4391-B904-6DAF7DC6D4BA;318n0EqLT3aRa4RQzznjIw
32839206-8FEA-45E3-B03C-953A432D3D1B;0oWv86Z_f5ux0ybJf3BJqR
7D2A3843-0322-44E3-8309-11CE37129DB5;1zAZX30o94uuC94Sut4fsr
CBDBF3B4-3CB6-47B6-B00E-BAD9E7A26318;3Bs$EqFBP7jh0EkjddecCO
0CDFD173-91B9-4EDC-9F3C-6B8B9045F8F3;0Ctz5paRbEt9yyQukGHVZp
C501E098-FBF5-48B2-A2E6-7B3EB4D5DC29;350U2O_$L8igBcUpwqrTmf
F96CA438-8364-4986-954E-6D50D0DC58DF;3vRAGuWsH9XfLERL3Gt5ZV
53908358-CCA4-44FB-A077-B45394254796;1Ja8DOpAH4_w1tj5EK9KUM
35666EDC-95F9-44FF-BC43-9A033B02FC8A;0rPcxSbVb4$xn3cWCx0loA
A862DD8A-2EED-4B32-85F5-8E6F060D4EAD;2eOjsABkrBCeNrZcy63Kwj
CE241872-6BCB-4FBB-9E96-793C2D69A81A;3E91XoQyjFkvwMUJmjQQWQ
F5F27EAA-5F67-4F33-AD68-74829FB5EF71;3rydwgNsTFCwreT8AVjUzn
A4236A40-08C1-41A4-B5F6-7A6FE04BC769;2a8sf02C51fBNsUc$WIyTf
A7F3B9F6-F29C-4BB4-B696-CF75FECBCD56;2dyxdsyfnBjBQMptN_oyrM
3FFAA6D4-E626-4F1B-AC41-696C7AFC799A;0$_gRKvYPF6wn1QMnw$7cQ
EC988658-A910-473A-9CE6-54318800AB3A;3ic8POgH17EfpcL3680Aiw
532AAA9F-F478-45D0-A123-3C2265A873D8;1JAggVz7X5qA4ZF29bg7FO
AED8D820-577C-424E-A628-050A4686C4A2;2ksDWWLtn2JgOe1Gf6XiIY
E71E7886-D293-4C9D-98FF-D8087C4BA542;3d7dY6qfDCdPZ$s0XyIwL2
FAE7C2F3-3DEA-47F1-9D8F-B23E2553AD58;3wvyBpFUf7yPsFiZubKwrO
D7967C0D-BC0F-43C5-B4B1-3F0C862056D1;3NbdmDl0z3nRInFmo685RH
F356F27A-51FD-4D4C-9AB4-7337EE66775E;3pLl9wKVrDJ9gqSpVkPdTU
F30CF677-B6EC-44EA-9724-C3DB2AF84DD3;3p3FPtjkn4wfSamzig_4tJ
038DAF3A-62A6-433E-9B3E-DBDDD0D66635;03ZQywOgP3Ffi_sztGrcOr
6E83C5F0-EB87-46C4-9CEC-19E678D685E0;1kWyNmwuT6n9pi6UPureNW
7C2DE7EA-8066-421C-9869-B1D28C9866D6;1yBUVgW6P279XfiTACc6RM
7D1324A1-5A13-44C8-B44C-E213A0740F8F;1z4oIXMXD4oBHCuXEWT0_F
7F50651E-C149-4568-9161-94B652ECDA20;1$K6KUmKb5Q95XbBPIxDeW
3377B1B1-DEBF-4A61-8CC8-54A72AEEFAC0;0pTx6nthzAOOp8LASgxlh0
1A399115-C053-481B-8759-5752A27F3B4F;0QEP4Lm5D86uTPLrAYVpjF
2C22163D-07BA-4509-82A3-610FAAAE7FAA;0i8XOz1xf52OAZOG_ghd_g
F908118F-C80C-4041-8C8C-D43744A2A258;3v216Fo0n0GOoCr3T4eg9O
9E47657D-3169-4125-946C-E70FB3288FFB;2UHsLzCMb19PHivm_pA8$x
56EA2EED-931A-46EB-BB5E-F7E8CBFA57FD;1MwYxjanf6wxjUz_ZB_bVz
3BF6FAAE-C312-40D8-B15A-0AB0458D9763;0xzlgkmn90sB5Q2h15ZPTZ
DEC4A41D-952D-4D19-91D4-7493D942934A;3UnAGTbIrD6P7KT9FPGfDA
3B7B413D-3AC0-4AD9-92F3-971932B67837;0xUq4zEi1AsPBpbnaojdWt
3B4954DE-86A0-4CA2-99DA-3B6FCA225C8D;0xILJUXg1CefdQEs$A8boD
AB76A22A-262C-4CD1-984C-55F203E47ECE;2hTg8g9YnCqPXCLV83v7xE
4614DE9B-0839-4992-B89B-F06B3A93611F;165DwR23b9ahYRy6iwas4V
B29D2E4D-9209-4EF1-AA55-9DF70BF727FE;2odIvDaWbEyQfLdVSBzoV_
`

	scanner := bufio.NewScanner(strings.NewReader(testData))
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		parts := strings.SplitN(line, ";", 2)
		if len(parts) != 2 {
			continue
		}

		guid := uuid.MustParse(parts[0])
		ifcGuid := parts[1]

		gotUuid, err := ToUuid(ifcGuid)
		assert.NoError(t, err)
		assert.Equal(t, guid, gotUuid)

		gotIfcGuid, err := FromUuid(gotUuid)
		assert.NoError(t, err)
		assert.Equal(t, ifcGuid, gotIfcGuid)
	}
}

func Test_New_and_ToUuid_and_FromUuid(t *testing.T) {
	for i := 0; i < 1000; i++ {
		ifcGuid, err := New()
		assert.NoError(t, err)
		assert.Len(t, ifcGuid, 22) // IFC GUID should always be 22 characters

		gotUuid, err := ToUuid(ifcGuid)
		assert.NoError(t, err)

		gotIfcGuid, err := FromUuid(gotUuid)
		assert.NoError(t, err)
		assert.Equal(t, ifcGuid, gotIfcGuid)

		gotU, err := ToUuid(gotIfcGuid)
		assert.NoError(t, err)
		assert.Equal(t, gotUuid, gotU)
	}
}

func Test_ConversionFunctions_with_invalid_data(t *testing.T) {
	tests := []struct {
		name    string
		ifcGuid string
		wantErr string
	}{
		{
			name:    "Empty string",
			ifcGuid: "",
			wantErr: "the ifcGuid must be 22 characters long",
		},
		{
			name:    "Too short",
			ifcGuid: "123456789012345678901",
			wantErr: "the ifcGuid must be 22 characters long",
		},
		{
			name:    "Too long",
			ifcGuid: "1234567890123456789012345",
			wantErr: "the ifcGuid must be 22 characters long",
		},
		{
			name:    "Invalid first character (greater than 3)",
			ifcGuid: "4ABCDEFGHIJKLMNOPQRSTU",
			wantErr: "illegal GUID '4ABCDEFGHIJKLMNOPQRSTU' found, it is greater than 128 bits",
		},
		{
			name:    "Invalid characters",
			ifcGuid: "ABC!@#$%^&*()_+{}|:<>?",
			wantErr: "contains invalid characters",
		},
		{
			name:    "All zeros",
			ifcGuid: "0000000000000000000000",
			wantErr: "",
		},
	}

	conversionFunctions := []struct {
		name     string
		function func(string) (any, error)
	}{
		{"ToUuid", func(s string) (any, error) { return ToUuid(s) }},
		{"ToInt64", func(s string) (any, error) { return ToInt64(s) }},
		{"ToInt32", func(s string) (any, error) { return ToInt32(s) }},
		{"ToIntString", func(s string) (any, error) { return ToIntString(s) }},
		{"ToAutoCadHandle", func(s string) (any, error) { return ToAutoCadHandle(s) }},
		{"ToString", func(s string) (any, error) { return ToString(s) }},
	}

	for _, tt := range tests {
		for _, cf := range conversionFunctions {
			t.Run(fmt.Sprintf("%s_%s", cf.name, tt.name), func(t *testing.T) {
				_, err := cf.function(tt.ifcGuid)
				if tt.wantErr != "" {
					assert.Error(t, err)
					assert.Contains(t, err.Error(), tt.wantErr)
				} else {
					assert.NoError(t, err)
				}
			})
		}
	}
}

func Test_FromUuid_with_uuidNil(t *testing.T) {
	ifcGuid, err := FromUuid(uuid.Nil)
	assert.Error(t, err)
	assert.Empty(t, ifcGuid)
	assert.Equal(t, err.Error(), "invalid UUID: nil UUID")
}

func Test_FromRevitUniqueId_with_valid_data(t *testing.T) {
	testData := `
00bdada5-6a16-4460-a1ce-b6ce6dc1cf00-001e72bd,00lQsbQXP4OA7Ejivjtxsz
00bdada5-6a16-4460-a1ce-b6ce6dc1cf00-001e75c6,00lQsbQXP4OA7Ejivjtxh6
00bdada5-6a16-4460-a1ce-b6ce6dc1cf00-001e75f4,00lQsbQXP4OA7Ejivjtxhq
00bdada5-6a16-4460-a1ce-b6ce6dc1cf00-001e75f8,00lQsbQXP4OA7Ejivjtxhu
03557dd0-b25d-459d-a58c-13ec43fb36d1-0023260d,03LNtGibr5dQMC4_n3s13S
05de7027-0d8f-47ba-b793-51d0941ed4ee-0026f65a,05td0d3Oz7khUJKT2KE2Aq
05de7027-0d8f-47ba-b793-51d0941ed4ee-0026f679,05td0d3Oz7khUJKT2KE2AN
05de7027-0d8f-47ba-b793-51d0941ed4ee-0026f68b,05td0d3Oz7khUJKT2KE29b
05de7027-0d8f-47ba-b793-51d0941ed4ee-0026f68f,05td0d3Oz7khUJKT2KE29X
05de7027-0d8f-47ba-b793-51d0941ed4ee-0026f697,05td0d3Oz7khUJKT2KE29v
05de7027-0d8f-47ba-b793-51d0941ed4ee-0026f773,05td0d3Oz7khUJKT2KE2ET
05de7027-0d8f-47ba-b793-51d0941ed4ee-0026f77f,05td0d3Oz7khUJKT2KE2EH
05de7027-0d8f-47ba-b793-51d0941ed4ee-0026f7f7,05td0d3Oz7khUJKT2KE2CP
05de7027-0d8f-47ba-b793-51d0941ed4ee-0026fbad,05td0d3Oz7khUJKT2KE2z3
05de7027-0d8f-47ba-b793-51d0941ed4ee-0026fc06,05td0d3Oz7khUJKT2KE2Ze
05de7027-0d8f-47ba-b793-51d0941ed4ee-0026fd62,05td0d3Oz7khUJKT2KE2cC
05de7027-0d8f-47ba-b793-51d0941ed4ee-0026fd90,05td0d3Oz7khUJKT2KE2b_
05de7027-0d8f-47ba-b793-51d0941ed4ee-0026fe0f,05td0d3Oz7khUJKT2KE2hX
05de7027-0d8f-47ba-b793-51d0941ed4ee-0026fe1d,05td0d3Oz7khUJKT2KE2hp
05de7027-0d8f-47ba-b793-51d0941ed4ee-0026ff06,05td0d3Oz7khUJKT2KE2le
05de7027-0d8f-47ba-b793-51d0941ed4ee-0026ff08,05td0d3Oz7khUJKT2KE2lc
05de7027-0d8f-47ba-b793-51d0941ed4ee-0026ff0a,05td0d3Oz7khUJKT2KE2la
05de7027-0d8f-47ba-b793-51d0941ed4ee-0026ff0c,05td0d3Oz7khUJKT2KE2lY
`

	scanner := bufio.NewScanner(strings.NewReader(testData))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}

		revitUniqueId := parts[0]
		ifcGuid := parts[1]

		gotIfcGuid, err := FromRevitUniqueId(revitUniqueId)
		assert.NoError(t, err)
		assert.Equal(t, ifcGuid, gotIfcGuid)
	}
}

func Test_FromRevitUniqueId_with_invalid_data(t *testing.T) {
	tests := []struct {
		name     string
		uniqueId string
	}{
		{
			name:     "Invalid length",
			uniqueId: "8d814f39-b6ea-4766-9a4f-8ac3de3501b2",
		},
		{
			name:     "Invalid format",
			uniqueId: "8d814f39-b6ea-4766-9a4f-8ac3de3501b2-XXXXXXXX",
		},
		{
			name:     "Empty string",
			uniqueId: "",
		},
		{
			name:     "Invalid characters",
			uniqueId: "8d814f39-b6ea-4766-9a4f-8ac3de3501b2-00007c0g",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromRevitUniqueId(tt.uniqueId)
			assert.Error(t, err)
			assert.Empty(t, got)
		})
	}
}

func Test_AutoCadHandle_conversions(t *testing.T) {
	tests := []struct {
		name      string
		handle    string
		wantError bool
	}{
		// Valid cases
		{
			name:      "Small number",
			handle:    "1A",
			wantError: false,
		},
		{
			name:      "Medium number",
			handle:    "DEADBEEF",
			wantError: false,
		},
		{
			name:      "Large number",
			handle:    "FFFFFFFFFFFFFFF",
			wantError: false,
		},
		{
			name:      "Max int64",
			handle:    "7FFFFFFFFFFFFFFF",
			wantError: false,
		},
		{
			name:      "Leading zeros",
			handle:    "000000ABCDEF",
			wantError: false,
		},
		{
			name:      "Mixed case",
			handle:    "aBcDeF123456",
			wantError: false,
		},
		// Invalid cases
		{
			name:      "Zero",
			handle:    "0",
			wantError: true,
		},
		{
			name:      "Empty string",
			handle:    "",
			wantError: true,
		},
		{
			name:      "Non-hexadecimal characters",
			handle:    "ABCDEFG",
			wantError: true,
		},
		{
			name:      "Floating point number",
			handle:    "1A.5",
			wantError: true,
		},
		{
			name:      "Exceeds max uint64",
			handle:    "80000000000000000",
			wantError: true,
		},
		{
			name:      "Contains spaces",
			handle:    "SOME SPACE",
			wantError: true,
		},
		{
			name:      "Contains prefix",
			handle:    "0xWithPrefix",
			wantError: true,
		},
		{
			name:      "Contains suffix",
			handle:    "WithSuffix_",
			wantError: true,
		},
		{
			name:      "Contains invalid characters",
			handle:    "Invalid-Characters",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIfcGuid, err := FromAutoCadHandle(tt.handle)

			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				gotHandle, err := ToAutoCadHandle(gotIfcGuid)
				assert.NoError(t, err)
				var wantHandle string
				if tt.handle == "0" {
					// if tt.handle == 0, gotHandle should also be 0
					wantHandle = tt.handle
				} else {
					// if tt.handle starts with zeros, gotHandle will not contain leading zeros
					wantHandle = strings.TrimLeft(tt.handle, "0")
				}
				assert.Equal(t, strings.ToLower(wantHandle), strings.ToLower(gotHandle))

				// Additional check: convert back to UUID and then to handle again
				gotUuid, err := ToUuid(gotIfcGuid)
				assert.NoError(t, err)

				finalHandle, err := uuidToAutoCadHandle(gotUuid)
				assert.NoError(t, err)
				assert.Equal(t, strings.ToLower(gotHandle), strings.ToLower(finalHandle))
			}
		})
	}
}

func Test_IntString_conversions(t *testing.T) {
	tests := []struct {
		name      string
		elementId string
		wantError bool
	}{
		// Valid cases
		{
			name:      "Small positive number",
			elementId: "123456789",
			wantError: false,
		},
		{
			name:      "Large positive number",
			elementId: "9223372036854775807", // Max int64
			wantError: false,
		},
		{
			name:      "Negative number",
			elementId: "-123456789",
			wantError: false,
		},
		// Invalid cases
		{
			name:      "Zero",
			elementId: "0",
			wantError: true,
		},
		{
			name:      "Empty string",
			elementId: "",
			wantError: true,
		},
		{
			name:      "Non-numeric string",
			elementId: "abc123",
			wantError: true,
		},
		{
			name:      "Floating point number",
			elementId: "123.456",
			wantError: true,
		},
		{
			name:      "Number too large for int64",
			elementId: "9223372036854775808", // Max int64 + 1
			wantError: true,
		},
		{
			name:      "Number too small for int64",
			elementId: "-9223372036854775809", // Min int64 - 1
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIfcGuid, err := FromIntString(tt.elementId)

			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				gotElementId, err := ToIntString(gotIfcGuid)
				assert.NoError(t, err)
				assert.Equal(t, tt.elementId, gotElementId)

				// Additional check: convert back to UUID and then to IntString again
				gotUuid, err := ToUuid(gotIfcGuid)
				assert.NoError(t, err)

				finalIntString, err := uuidToIntString(gotUuid, "%d")
				assert.NoError(t, err)
				assert.Equal(t, tt.elementId, finalIntString)
			}
		})
	}
}

func Test_Int32_conversions(t *testing.T) {
	tests := []struct {
		name      string
		elementId int32
		wantError bool
	}{
		{
			name:      "Small positive number",
			elementId: 123456789,
		},
		{
			name:      "Maximum int32",
			elementId: math.MaxInt32,
		},
		{
			name:      "Minimum int32",
			elementId: math.MinInt32,
		},
		{
			name:      "Negative number",
			elementId: -123456789,
		},
		{
			name:      "Zero",
			elementId: 0,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIfcGuid, err := FromInt32(tt.elementId)
			if tt.wantError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			gotElementId, err := ToInt32(gotIfcGuid)
			assert.NoError(t, err)
			assert.Equal(t, tt.elementId, gotElementId)

			// Additional check: convert back to UUID and then to Int32 again
			gotUuid, err := ToUuid(gotIfcGuid)
			assert.NoError(t, err)

			finalInt32, err := uuidToInt64(gotUuid)
			assert.NoError(t, err)
			assert.Equal(t, tt.elementId, int32(finalInt32))
		})
	}
}

func Test_Int64_conversions(t *testing.T) {
	tests := []struct {
		name      string
		objectId  int64
		wantError bool
	}{
		{
			name:     "Maximum int64",
			objectId: math.MaxInt64,
		},
		{
			name:     "Maximum int32",
			objectId: math.MaxInt32,
		},
		{
			name:     "Minimum int32",
			objectId: math.MinInt32,
		},
		{
			name:     "Positive number",
			objectId: 123456789,
		},
		{
			name:     "Large number",
			objectId: 9223372036854775807,
		},
		{
			name:     "Negative number",
			objectId: -123456789,
		},
		{
			name:      "Zero",
			objectId:  0,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIfcGuid, err := FromInt64(tt.objectId)
			if tt.wantError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			gotObjectId, err := ToInt64(gotIfcGuid)
			assert.NoError(t, err)
			assert.Equal(t, tt.objectId, gotObjectId)

			// Additional check: convert back to UUID and then to Int64 again
			gotUuid, err := ToUuid(gotIfcGuid)
			assert.NoError(t, err)

			finalInt64, err := uuidToInt64(gotUuid)
			assert.NoError(t, err)
			assert.Equal(t, tt.objectId, finalInt64)
		})
	}
}

func Test_String_conversions(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantError bool
	}{
		// Valid cases
		{
			name:      "Revit file with ID",
			input:     "project1.rvt|123456",
			wantError: false,
		},
		{
			name:      "AutoCAD file with handle",
			input:     "drawing.dwg|3A5C",
			wantError: false,
		},
		{
			name:      "Civil3D file with ID",
			input:     "site_plan.dwg|1A2B3C",
			wantError: false,
		},
		{
			name:      "Alphanumeric string",
			input:     "ABC123xyz789",
			wantError: false,
		},
		{
			name:      "String with special characters",
			input:     "Project_2023-05-15@Rev2",
			wantError: false,
		},
		{
			name:      "Short string",
			input:     "abc",
			wantError: false,
		},
		{
			name:      "Long string",
			input:     "ThisIsAVeryLongStringThatExceedsTheTypicalLengthOfAnIdentifierButShouldStillBeValid",
			wantError: false,
		},
		{
			name:      "String with leading and trailing spaces",
			input:     "  ABC123xyz789  ",
			wantError: false,
		},
		{
			name:      "Zero",
			input:     "0",
			wantError: false,
		},
		{
			name:      "String with unicode characters",
			input:     "��ÖÜäöüß",
			wantError: false,
		},
		// Invalid cases
		{
			name:      "Empty string",
			input:     "",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test FromString
			gotIfcGuid, err := FromString(tt.input)

			if tt.wantError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Len(t, gotIfcGuid, 22) // IFC GUID should always be 22 characters

			// Test ToString
			gotString, err := ToString(gotIfcGuid)
			assert.NoError(t, err)
			t.Logf("Input string: `%s`", tt.input)
			t.Logf("Got string: `%s`", gotString)

			// The original input and the result of ToString may not be identical.
			// We can't directly compare them.
			// Instead, we ensure that FromString produces the same result for both.
			checkIfcGuid, err := FromString(gotString)
			assert.NoError(t, err)
			assert.Equal(t, gotIfcGuid, checkIfcGuid)

			// Additional check: convert IFC GUID to UUID and back
			gotUuid, err := ToUuid(gotIfcGuid)
			assert.NoError(t, err)

			finalIfcGuid, err := FromUuid(gotUuid)
			assert.NoError(t, err)
			assert.Equal(t, gotIfcGuid, finalIfcGuid)
		})
	}
}
