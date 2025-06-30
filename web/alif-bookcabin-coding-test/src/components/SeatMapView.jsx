import React, { useEffect, useState } from "react";
import { fetchSeatMap } from "../api/seatMapAPI";
import SeatGrid from "./SeatGrid";
import PassengerDetails from "./PassengerDetails";
import SegmentDetails from "./SegmentDetails";

const SeatMapView = ({ itineraryId }) => {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);

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

  const segment =
    data?.data?.seatsItineraryParts?.[0]?.segmentSeatMaps?.[0]?.segment;

  const passenger =
    data?.data?.seatsItineraryParts?.[0]?.segmentSeatMaps?.[0]
      ?.passengerSeatMaps?.[0]?.passenger;

  return (
    <div
      style={{
        display: "flex",
        gap: "24px",
        alignItems: "flex-start",
      }}
    >
      <div style={{ flex: 1 }}>
        <SegmentDetails segment={segment} />
      </div>
      <div style={{ flex: 2 }}>
        <h2 style={{ textAlign: "center" }}>Seat View</h2>
        <SeatGrid seatMapData={data} />
      </div>
      <div style={{ flex: 1 }}>
        <PassengerDetails passenger={passenger} />
      </div>
    </div>
  );
};

export default SeatMapView;
