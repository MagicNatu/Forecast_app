# Weather monitoring application
Provides current and forecast weather for N locations

On windows, compile by running RunMe.bat (go must be installed)

Alternatively, in the working directory run the following commands:
//
Go build
./.exe
//

To access the application UI, navigate to "localhost:8000/main".

JSON data is available from within the application, or by navigating to "localhost:8000/forecast/{add location}/{number_of_days}".
Forecasts are only available for number_of_days < 5.

To run associated tests, run "go test -v"

******
For those who are unable to compile, there is a pre-compiled binary provided
******
