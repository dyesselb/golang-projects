# groupie-tracker-search-bar

### Description

Groupie Trackers consists on receiving a given API and manipulate the data contained in it, in order to create a site, displaying the information.

- It was given an [API](https://groupietrackers.herokuapp.com/api), that consists in four parts:

  - The first one, `artists`, containing information about some bands and artists like their name(s), image, in which year they began their activity, the date of their first album and the members.

  - The second one, `locations`, consists in their last and/or upcoming concert locations.

  - The third one, `dates`, consists in their last and/or upcoming concert dates.

  - And the last one, `relation`, does the link between all the other parts, `artists`, `dates` and `locations`.

- Given all this a user friendly website was built where you can see the bands info through several data visualizations.

The search bar has typing suggestions as you type. 

- Following search cases are handled:  
  > 
  > artist/band name   
  >  
  > members    
  > 
  > locations    
  >
  > first album date   
  >
  > creation date   
  >

- This project also focuses on the creation of events/actions and on their visualization.


### Instructions

- To run the program write: `$go run cmd/main.go`

- Open `http://localhost:4000`


### This project implements :

- Manipulation and storage of data.
- [JSON](https://www.json.org/json-en.html) files and format.
- HTML.
- Event creation and display.
- [Client-server](https://developer.mozilla.org/en-US/docs/Learn/Server-side/First_steps/Client-Server_overview).

### Useful links

- https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/04.1.html
- https://www.golangprograms.com/goroutines-and-channels-example.html
