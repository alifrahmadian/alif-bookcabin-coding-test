import React, { useState } from "react";

const SeatBox = ({ seat }) => {
  const [show, setShow] = useState(false);

  return (
    <div
      onMouseEnter={() => setShow(true)}
      onMouseLeave={() => setShow(false)}
      style={{
        position: "relative",
        width: 30,
        height: 30,
        margin: 2,
        backgroundColor: seat.available ? "#4caf50" : "#ccc",
        display: "inline-block",
        textAlign: "center",
        lineHeight: "30px",
        cursor: "pointer",
      }}
    >
      {seat.code || "-"}

      {show && (
        <div
          style={{
            position: "absolute",
            top: "-110px",
            left: "-50px",
            width: "150px",
            background: "#fff",
            border: "1px solid #ccc",
            padding: "8px",
            fontSize: "12px",
            zIndex: 100,
            boxShadow: "0 2px 5px rgba(0,0,0,0.3)",
          }}
        >
          <div>
            <strong>Code:</strong> {seat.code || "-"}
          </div>
          <div>
            <strong>Status:</strong>{" "}
            {seat.available ? "Available" : "Not Available"}
          </div>
          <div>
            <strong>Price:</strong>{" "}
            {seat.total?.alternatives?.[0]?.[0]?.amount ?? "Free"}{" "}
            {seat.total?.alternatives?.[0]?.[0]?.currency ?? ""}
          </div>
          <div>
            <strong>Refund:</strong> {seat.refundIndicator || "-"}
          </div>
          <div>
            <strong>Characteristics:</strong>{" "}
            {seat.seatCharacteristics?.join(", ") || "-"}
          </div>
        </div>
      )}
    </div>
  );
};

export default SeatBox;
