import React from "react";

const PassengerDetails = ({ passenger }) => {
  if (!passenger) return null;

  const { passengerDetails, passengerInfo, preferences, documentInfo } =
    passenger;

  return (
    <div
      style={{
        border: "1px solid #ccc",
        borderRadius: "4px",
        padding: "12px",
        maxWidth: "400px",
        fontSize: "14px",
      }}
    >
      <h3>
        {passengerDetails?.firstName || "-"} {passengerDetails?.lastName || ""}
      </h3>
      <div>
        <strong>Passenger Type:</strong> {passengerInfo?.type || "-"}
      </div>
      <div>
        <strong>Gender:</strong> {passengerInfo?.gender || "-"}
      </div>
      <div>
        <strong>Date of Birth:</strong> {passengerInfo?.dateOfBirth || "-"}
      </div>
      <div>
        <strong>Email:</strong>{" "}
        {passengerInfo?.emails?.length > 0
          ? passengerInfo.emails.join(", ")
          : "-"}
      </div>
      <div>
        <strong>Phone:</strong>{" "}
        {passengerInfo?.phones?.length > 0
          ? passengerInfo.phones.join(", ")
          : "-"}
      </div>
      <div style={{ marginTop: "8px" }}>
        <strong>Address:</strong>
        <div>
          {passengerInfo?.address?.street1 || "-"},{" "}
          {passengerInfo?.address?.street2 || "-"}
        </div>
        <div>
          {passengerInfo?.address?.city || "-"},{" "}
          {passengerInfo?.address?.state || "-"}{" "}
          {passengerInfo?.address?.postcode || "-"}
        </div>
        <div>{passengerInfo?.address?.country || "-"}</div>
      </div>
      {preferences?.frequentFlyer?.length > 0 && (
        <div style={{ marginTop: "8px" }}>
          <strong>Frequent Flyer:</strong>
          {preferences.frequentFlyer.map((ff, idx) => (
            <div key={idx}>
              {ff.airline} - {ff.number} (Tier {ff.tierNumber})
            </div>
          ))}
        </div>
      )}
      {documentInfo && (
        <div style={{ marginTop: "8px" }}>
          <strong>Document Info:</strong>
          <div>
            {documentInfo.documentType || "-"} -{" "}
            {documentInfo.nationality || "-"}
          </div>
        </div>
      )}
    </div>
  );
};

export default PassengerDetails;
