import React from "react";
import SeatBox from "./SeatBox";

const SeatGrid = ({ seatMapData }) => {
  const segments =
    seatMapData?.data?.seatsItineraryParts?.[0]?.segmentSeatMaps || [];

  return (
    <div>
      {segments.map((segment, idx) => (
        <div key={idx}>
          {segment.passengerSeatMaps.map((psm, pIdx) => (
            <div key={pIdx}>
              {psm.seatMap.cabins.map((cabin, cIdx) => (
                <div key={cIdx}>
                  {cabin.seatRows.map((row, rowIdx) => (
                    <div key={rowIdx} style={{ display: "flex" }}>
                      {(row?.seats || []).map((seat, sIdx) =>
                        seat ? (
                          <SeatBox key={sIdx} seat={seat} />
                        ) : (
                          <div
                            key={sIdx}
                            style={{
                              width: 30,
                              height: 30,
                              background: "#eee",
                            }}
                          >
                            Empty
                          </div>
                        )
                      )}
                    </div>
                  ))}
                </div>
              ))}
            </div>
          ))}
        </div>
      ))}
    </div>
  );
};

export default SeatGrid;
