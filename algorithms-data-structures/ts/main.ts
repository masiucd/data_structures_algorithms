const isUniqueTicketNumber = (tabs, ticketNumber) => {
  return (
    tabs.filter((tab) => tab.selectedTicketNumber === ticketNumber).length <= 1
  );
};
