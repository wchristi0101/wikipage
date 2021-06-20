# wikipage
Given a wiki page URL, gives data related to it

Code is generated from OpenAPI 3.0. You can find documentation in the documentation folder or at https://app.swaggerhub.com/apis/VoithAPI/WikipediaAPI/1.0.1

To run the server goto ~/go-server-server-generated and run go run main.go

To access data from the server use information found in the documentation. 

The base function is http://localhost:8080/VoithAPI/WikipediaAPI/1.0.1/wikipages?url=https://wikipedia.org/wiki/Heidenheim_an_der_Brenz

You can sort this with a filter=["filtera","filterb"]

If given time it would have been easy to reuse the same functions that generated the base data. Since DRY principles were used, all other functions could have been sorted, and filters could have easily been applied to them. There was not enough time to complete this, however.

I chose to build this in go. I spent more time than allotted, because I am just starting to learn go, and I figured this was a good project to learn more about it. I am certain there are more efficient ways to do this, and if building in Java or Python the whole project likely would have been completed. The algorithms I would have used to solve the problems are commented in the code.

The most important folder is ~/go-server-server-generated/go/api_default.go. This is where the business logic is at. 
