import "./App.css";
import SeatMapView from "./components/SeatMapView";

function App() {
  return (
    <div>
      <h1>Bookcabin Coding Test</h1>
      <SeatMapView itineraryId={1} />
    </div>
  );
}

export default App;
