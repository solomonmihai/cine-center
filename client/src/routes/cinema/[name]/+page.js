import apiCall from "$lib/services/Api";

export async function load({ params }) {
  const { name } = params;
  return await apiCall(`/cinema/${name}`);
}
