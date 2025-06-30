import React, { useEffect, useState } from "react";
import { fetchSeatMap } from "../api/seatMapAPI";

const SeatMapView = ({ itineraryId }) => {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(null);

  useEffect(() => {
    const loadData = async () => {
      try {
        const seatMap = await fetchSeatMap(itineraryId);
        setData(seatMap);
      } catch (error) {
        console.error("Error fetching seat map:", error);
      } finally {
        setLoading(false);
      }
    };

    loadData();
  }, [itineraryId]);

  if (loading) return <div>Loading seat map...</div>;

  return (
    <div>
      <h2>Seat Map</h2>
      <p>{JSON.stringify(data, null, 2)}</p>
    </div>
  );
};

export default SeatMapView;
