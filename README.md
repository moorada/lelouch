# lelouch
The description of this tool is in Italian because this tool is designed for Italians (or for those who know Italian).

Tool per esercitarsi con la conversione fonetica (The major system), in lingua italiana.

## 🕹️ Il Mio Gioco - Versione CLI  

🚀 Questa è la vecchia versione del mio gioco, disponibile solo da riga di comando.  

🔥 **Ora è disponibile una versione web migliorata!** 🔥  

📜 La nuova versione web offre un'interfaccia più intuitiva e una migliore esperienza di gioco.  

👉 Gioca online qui: [VERSIONE WEB](https://moorada.github.io/Converti)  


## Conversione fonetica
La conversione fonetica è una tecnica di memorizzazione dei numeri. Funziona convertendo i numeri in consonanti e, aggiungendo opportunamente delle vocali, trasformarle in parole che si possono ricordare con più facilità di una serie di numeri, in modo particolare usando altre regole mnemoniche. 

Ogni cifra viene convertita in una consonante. 
| Num. |	Suono |	Lettere |	Esempio
| --- | --- | --- | --- |
1 |	dentale |	T, D | 	tè, dio, ateo, due, atto |
2 |	nasale |	N, GN |	neo, anno, gnè |
3 |	mugolante |	M | amo, mio, emme |
4 |	vibrante |	R |	ara, re, oro, erre |
5 |	liquido |	L, GL |	ali, lui, aglio, li |
6 |	palatale |	C, G (dolci) |	ciao, oggi, ci, gi, agio |
7 |	gutturale |	C, G (dure), K |	occhio, eco, chi, qui, ago, gay, acca |
8 |	labiodentale |	F, V |	ufo, uva, via, uffa, avvio |
9 |	labiale |	P, B |	boa, ape, oppio, oboe |
0 |	sibilante |	S, SC, Z |	sei, esse, zio, ozio, ascia, scia |


Per convertire una parola in un corrispondente numerico (e viceversa) vanno rispettate alcune regole:

* Le vocali non corrispondono a nessuna cifra, quindi non vanno considerate;
* Le consonanti doppie vanno considerate come un unico suono;
* Bisogna sempre valutare il suono che la lettera produce. Per esempio, "gl" in "sciogliere" ha un suono liquido, perciò corrisponde a 5, mentre in "glicine" "gl" produce due suoni separati ("g" gutturale e "l" liquida), corrispondenti a 75. Sempre in "sciogliere", "sc" produce un suono sibilante corrispondente a 0, mentre in "scatola" "sc" produce due suoni separati ("s" sibilante e "c" gutturale), corrispondenti a 07.

Grazie alla conversione fonetica è possibile ricordare una quantità innumerevole di cifre. Di seguito un esempio per ricordare il Pi greco fino alla trentaduesima cifra decimale.
* Pi greco con 32 cifre decimali: 3,14159265358979323846264338327950
* Una TROTA ALPINA voleva volare fino in CIELO, ma prima di partire si mise la MAGLIA, perché aveva paura del freddo: una vera FOBIA. Arrivata in quota incontrò un'OCA, dalla cui coda mancavano delle PIUME. Gliele aveva strappate uno GNOMO VORACE, che quando non mangia oche si sazia divorando NOCI, noci che coglie dai RAMI coperti di MUFFA, sporcandosi la MANICA vicino al POLSO. 

```
(3,) 
TROTA   141
ALPINA  592
CIELO    65
MAGLIA   35
FOBIA    89
OCA       7
PIUME    93
GNOMO    23 
VORACE  846
NOCI     26
RAMI     43
MUFFA    38
MANICA  327
POLSO   950
```

## Funzionalità
<p align="center">
  <img alt="Interfaccia Lelouch" src="https://raw.githubusercontent.com/moorada/lelouch/master/img/menu.png" width="40%" />
</p>

"Converti numeri o parole", converte semplicemente numeri/parole in parole/numeri. Nella conversione in parole, vengono visualizzate tutte le possibili combinazioni (delle parole comuni) in una tabella. Es:
<p align="center">
  <img alt="Interfaccia Lelouch" src="https://raw.githubusercontent.com/moorada/lelouch/master/img/convertitore.png" width="90%" />
</p>

Le modalità "Gioco", interrogano, controllano la risposta e suggeriscono altre soluzioni.
<p align="center">
  <img alt="Interfaccia Lelouch" src="https://raw.githubusercontent.com/moorada/lelouch/master/img/convertiParole.png" width="40%" />
</p>

Inoltre è possibile scegliere il livello
<p align="center">
  <img alt="Interfaccia Lelouch" src="https://raw.githubusercontent.com/moorada/lelouch/master/img/sceltaLivello.png" width="40%" />
</p>

Se si vuole aggiungere o rimuovere delle parole al dizionario, basta modificare il file "dizionari/parolecomuni.txt", ed eliminare i file *.json. All'avvio del tool i dizionari numerati verranno automaticamente ricreati.
