# lelouch
Tool per esercitarsi con la conversione fonetica (The major system), in lingua italiana.

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

## Funzionalità
<p align="center">
  <img alt="Interfaccia Lelouch" src="https://raw.githubusercontent.com/moorada/lelouch/master/img/interfaccia.png" width="40%" />
</p>

* La prima funzione crea un dizionario numerato (ogni parola è associata al numero corrispondente della conversione fonetica) con tutte le parole presenti nei file .txt nella cartella "paroleitaliane".
* "CONVERTI", converte semplicemente il numero/parola in parola/numero.
* "GAME", interroga e controlla la tua risposta.

## In sviluppo
* Scegliere la difficoltà del gioco, selezionando l'intervallo della lunghezza delle parole/numeri.
* Modalità di gioco con più parole.
* Modalità di gioco personale con numeri/parole/frasi inserite dall'utente.
