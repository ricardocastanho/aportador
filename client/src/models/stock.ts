export type Stock = {
  ticker: string;
  price: string;
  shares: string;
  profit: string;
  payout: string;
  dpa: number;
  ceilPrice: number;
};

export function calculateDpa(profit: string, shares: string, payout: number): number {
  const lpa =
    parseAmountToNumber(profit) / parseAmountToNumber(shares);

  const dpa = lpa * (Number(payout) / 100);

  return dpa;
}

export function getCeilPrice(dpa: number, dy: number): number {
  const ceilPrice = (dpa / dy) * 100;
  return ceilPrice;
}

export function parseAmountToNumber(amount: string): number {
  const formatted = amount.replaceAll(",", ".");
  const splitted = formatted.split(".");

  if (splitted.length > 3) {
    return Number(`${splitted[0]}.${splitted[1]}`);
  }

  if (splitted.length > 2) {
    return Number(`0.${splitted[0]}${splitted[1]}`);
  }

  return Number(splitted.join("."));
}
