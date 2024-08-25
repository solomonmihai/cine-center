import apiCall from "$lib/services/Api";

export async function load() {
  const cinemas = await apiCall("/cinema-names");
  const films = await apiCall("/all-films");

  return { cinemas, films };
}
