# Weather monitoring application
Provides current weather conditions and forecast predictions for N locations

On windows, compile by running RunMe.bat (go must be installed)

Alternatively, in the working directory run the following commands:
//
Go build
./.exe
//

To access the application UI, navigate to "localhost:8000/main".

JSON data is available from links within the application. It is also possible to add locations from an endpoint, by navigating to "localhost:8000/forecast/{add_location}/{days}/{min_temp}/{max_temp}"

Forecast data is available for days <= 5.

It is possible to retrieve data from "localhost:8000/JSONforecast"

To run associated tests, run "go test -v"

******
For those who are unable to compile, there is a pre-compiled binary provided
******
