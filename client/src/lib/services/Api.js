const BASE_URL = "/api";
const HEADERS = {
  "Content-Type": "application/json",
};

/**
 * @param {any} url
 * @param {any} method
 * @param {any} data
 */
export default async function apiCall(url) {
  try {
    const res = await fetch(BASE_URL + url, {
      method: "GET",
      headers: HEADERS,
    });
    return res.json();
  } catch (err) {
    console.log(`failed to call api: ${url}, ${err}`);
  }
}
