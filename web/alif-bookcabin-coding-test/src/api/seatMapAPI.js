import axios from "axios";

const API_BASE_URL = "http://localhost:8080";

export const fetchSeatMap = async (seatsItineraryPartId) => {
  const response = await axios.get(
    `${API_BASE_URL}/seat-map/seats-itinerary-part/${seatsItineraryPartId}`
  );

  return response.data;
};
