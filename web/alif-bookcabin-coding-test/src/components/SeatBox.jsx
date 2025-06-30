import React from "react";

const SeatBox = ({ seat }) => {
  console.log("SeatBox:", seat);

  if (!seat) return <div>EMPTY</div>;
  return (
    <div
      title={`Seat ${seat.code || "-"} - ${
        seat.total?.alternatives?.[0]?.[0]?.amount || "Free"
      }`}
      style={
        seat.storefrontSlotCode === "SEAT"
          ? {
              width: 30,
              height: 30,
              margin: 2,
              backgroundColor: seat.available ? "#4caf50" : "#ccc",
              display: "inline-block",
              textAlign: "center",
              lineHeight: "30px",
              cursor: "pointer",
            }
          : {
              margin: 2,
              height: 30,
              width: 5,
              lineHeight: "30px",
              display: "inline-block",
              textAlign: "center",
              backgroundColor: "#ccc",
            }
      }
    >
      {seat.code}
    </div>
  );
};

export default SeatBox;
