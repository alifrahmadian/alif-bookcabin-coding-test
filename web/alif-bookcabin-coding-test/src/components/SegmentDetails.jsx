import React from "react";

// Utility function biar konsisten format
const formatDate = (isoString) => {
  if (!isoString) return "-";
  const date = new Date(isoString);
  return date.toLocaleString("en-US", {
    dateStyle: "medium",
    timeStyle: "short",
  });
};

const SegmentDetails = ({ segment }) => {
  if (!segment) return null;

  return (
    <div
      style={{
        border: "1px solid #ccc",
        borderRadius: "4px",
        padding: "12px",
        maxWidth: "300px",
        fontSize: "14px",
      }}
    >
      <h3>Segment Info</h3>
      <div>
        <strong>Origin:</strong> {segment.origin}
      </div>
      <div>
        <strong>Destination:</strong> {segment.destination}
      </div>
      <div>
        <strong>Departure:</strong> {formatDate(segment.departure)}
      </div>
      <div>
        <strong>Arrival:</strong> {formatDate(segment.arrival)}
      </div>
      <div>
        <strong>Departure Terminal:</strong>{" "}
        {segment.flight?.departureTerminal || "-"}
      </div>
      <div>
        <strong>Arrival Terminal:</strong>{" "}
        {segment.flight?.arrivalTerminal || "-"}
      </div>
      <div>
        <strong>Cabin Class:</strong> {segment.cabinClass}
      </div>
      <div>
        <strong>Flight Number:</strong> {segment.flight?.flightNumber}
      </div>
      <div>
        <strong>Airline Code:</strong> {segment.flight?.airlineCode}
      </div>
      <div>
        <strong>Equipment:</strong> {segment.equipment}
      </div>
      <div>
        <strong>Booking Class:</strong> {segment.bookingClass}
      </div>
      <div>
        <strong>Fare Basis:</strong> {segment.fareBasis}
      </div>
    </div>
  );
};

export default SegmentDetails;
