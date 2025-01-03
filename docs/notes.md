# Thoughts

- Laver en priorites ko, som er nem at skifte prioritest ko til
- hver element i ko'en holder en note, samt en under ko

- I parent elementet, skal man kunne se alle borns noter
    - maaske nice til at se hvor langt man er kommet?

- Skal man stadig kunne se noten, efter tasken er complete?

- nogen gange ville det vaere nice at have baade liste og noter ved siden af hindanden
    - Saa skal man ikke hoppe saadan frem og tilbage

- Skal det vaere en stack
    - naar siden loader, vil den overste task altid vaere den man ser forst <= current task 
        - hvordan skal man nemt bytte rundt paa prioriteter
        - keybinding til inc/dec prioritet
            - saa alle sub tasks ogsaa flytter med

- Byg det som et API, saa man kan bruge det fra baade web og som et nvim plugin.

- Skal vi bygge hjemmesiden saa den er nem at konvereter til at bruge api
    - skal man lave hjemmeside eller API forst
        - Man burde jo lave API'et forst, ellers har man bare en skal

- Saa skal jeg prove at decouple backenden, saa meget at det er nemt at "swappe" backend til et API naar det er klart

- https://www.reddit.com/r/golang/comments/1hahlgs/i_built_my_personal_website_completely_in_go/?share_id=eMqCrW1SjnWdBjovCoz9D&utm_name=iossmf
     - han har brugt golang, templ og htmx til at bygge hans website
     - Hans kodebase er meget clean

- lets go repo
    - https://github.com/DataDavD/letsgo_snippetbox?tab=readme-ov-file

## Plan

- webserver laves i golang 
- frontend laves i React

- Lav det med webassembly
    - https://permify.co/post/wasm-go/
    - https://thenewstack.io/webassembly-and-go-a-guide-to-getting-started-part-1/
    - https://medium.com/@joloiuy/go-beyond-the-browser-embracing-webassembly-with-go-ccc6d97e8b64
        - webassembly har ikke nogen garbage collector

## krav

- nemt at se hvilken task der er current
- nemt overblik over all tasks
- noter til hver task 
- inc/dec task prioritet med keypress
    - hvor alle subtasks folger med
    - skal gores hurtigt

## Indhold

- nogen gange vil man bare have et lille punkt, og en linje el 2 med noter og intet andet.  Hvor det ville være træls at skulle hoppe over i et helt nyt element

- note og tasklist skal gå begge veje, så hvis man laver et nyt punkt i noten, så kommer den automatisk på tasklist

- Det skal være nemt at add et nyt element til køen

- I en note skal man nemt kunne lave et underpunkt 
    - skal det hele styres fra note appen

## faseability -> lav hjemmeside for API

- normalt laver man hjemmesode til sidst. Men nu har jeg lige laest hvordan man bygger et i go, og saa ville det maaske vaere lidt spildt, hvis jeg kaster mig over neget andet

- hvordan vil man gore det, hvis man skulle ligge op til at bruge et API i fremtiden
    - naar man ikke bruger et API, vil man jo selv have alt sql kald osv i handlers. Det laver man allegivel et "mini" bibliotek til med select, insert og update.
    - Saa har men en task model i weben, som skal matche den i databasen. Og API'et skal bruge samme model
        - Kan man lave det saa API modellen maaske har flere felter end hvad der bruges paa web
        - saa skal man have interfaces
    
    - REST er bare lidt glorifyed sql procedures, saa tror ikke omvaeltningen er saa stor

- konkrete tiltag for API integration
    - Interface der definere alle metoderne og implementers af en struct
    - Saa kalder vi de samme metoder i handlers

## Webserver der modtager post, get med noter og opdatere en database 

### note model til at snakke med database 

### Opret database 

### Opsaetning af webserver

- Hvordan laver vi dependency injection, naar man ikke kan add nye metoder til vores app, hvis ikke det er i samme package

## react frontend

- frontend til webserver til at get og post noter til database, gennem browseren

## webassembly app til at korere det hele i browseren

- sikkerhed er en consern, fordi client downloader alt koden, saa man skal ikke have nogen hardcoded secrets

## dependency injection 

- Hvordan laver vi dependency injection i vores middleware 
    - Vi vil bruge logger fra vores dependency, 
    - middleware ligger i sin egen pakke, saa man kan ikke bare add en middleware i config pakken. 
    - middleware skal have en http.Handler signatur, som den ikke kan faa hvis vi injecter app som normalt.

- At bruge en struct til dependency injection kan heller ikke vaere den eneste maade at gore det paa.

- DI handler bare om at lave et interface der repraesentere den dependency du skal brug
- Man vil bare ikke binde sig haardt til en dependency, man vil skifte ud 
- Saa man laver et interface der repraesentere dependency'en og saa giver man en instans med naar man kalder 

- logging, behoves ikke at injectes.
    - Hvordan skal man haandtere logging


### Hvordan skal man haandtere logging

- hvis ikke vi skal inject det i vores app 

- kan lave en logger struct, med clientInfo og serverError messages

### Hvordan skal man inject database 
 
- vil man lave en interface, med alle db metoderne, eller vil man kun have vores connection pool.

### Hvordan injecter vi i vores handlers

- de forskellige dependencies skal ligge paa vores app 

- Det gor vi som i bogen.

- Vi skal lave en logger struct, i config.app. saa alle metoderne ligger paa logger structen. og ikke direkte koblet paa app.

### Hvordan injecter vi i vores middleware

- faa helt styr paa hvorfor det fungere 


### inject loggers

- app holder vores 3 log metoder

- naar man creater app, kan man inject sin egen logger 
- Hvis vi skal log til en fil, skal vi have en struct som matcher log.Logger. Hvilket nok bliver lidt svaeret 


- Men vi vil gerne have bruge vores egen logger formattering 
    - logging metoderne kan bare tage app, eller en Logger som input


