# GolangPraktischerTeil
Dieses Projekt ist ein kleines Programm (parallel.go), welches im Rahmen des Seminars "Webprogrammierung" an der Friedrich-Schiller-Universität Jena erstellt wurde und zur Veranschaulichung der Parallelität der Programmiersprache Go dienen soll.


## Go-Installation:

Um Go "zum laufen zu bekommen", muss man selbstverständlich [Go](https://golang.org) herunterladen und installieren. 
Weitere Informationen zum Installationsprozess findet man hier: https://golang.org/doc/install

## Zum Praktischen Teil:

Um zum Beispiel das Programm (parallel.go) auszuführen, reicht es auch auf den Go-eigenen [Spielplatz](https://play.golang.org) zu gehen, den Quelltext der Datei zu copy-and-pasten und ihn auszuführen. 
Das Programm selbst dient zur Veranschaulichung einer der wichtigsten Funktionen von Go, der Nebenläufigkeit. Wichtig bedeutet in diesem Fall, dass sich Go damit von anderen Programmiersprachen abhebt.
Verschiedene "worker pools" greifen dabei auf einen "channel" zu, und machen dann einzelne Dinge damit. 
