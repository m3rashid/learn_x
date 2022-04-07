import React from "react";
import ReactWeather, { useOpenWeather } from "react-open-weather";

function App() {
  const { data, isLoading, errorMessage } = useOpenWeather({
    key: "ee3f649a1f91a5553471f39abb45e300",
    lat: "1.2921",
    lon: "36.8219",
    lang: "en",
    unit: "metric", // values are (metric, standard, imperial)
  });

  return (
    <div className="App">
      <ReactWeather
        isLoading={isLoading}
        errorMessage={errorMessage}
        data={data}
        lang="en"
        locationLabel="Nairobi"
        unitsLabels={{ temperature: "C", windSpeed: "Km/h" }}
        showForecast
      />
    </div>
  );
}

export default App;
