# Forecast_app
Provides current and forecast weather for N locations

On windows, compile by running RunMe.bat (go must be installed)

Alternatively, in the working directory run the following commands:
//
Go build
./.exe
//

To access the application UI, navigate to "localhost:8000".

JSON data is available by navigating to "localhost:8000/current/{location}" or "localhost:8000/forecast/{location}/{days}"

To run associated tests, run "go test -v"

******
For those who don't have go installed, there is a pre-compiled binary provided
******
