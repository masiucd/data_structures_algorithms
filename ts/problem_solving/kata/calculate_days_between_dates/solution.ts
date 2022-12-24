export function calculateDaysBetweenDates(begin: string, end: string) {
  const beginDate = new Date(begin)
  const endDate = new Date(end)
  return (endDate.getTime() - beginDate.getTime()) / (1000 * 60 * 60 * 24)
}
