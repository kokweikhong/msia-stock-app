import { z } from "zod";

export const QuoteSchema = z.object({
  code: z.string(),
  name: z.string(),
  short_name: z.string(),
  category: z.string(),
  sector: z.string(),
  eps: z.number(),
  nta: z.number(),
  pe: z.number(),
  dy: z.number(),
  roe: z.number(),
});

export type Quote = z.infer<typeof QuoteSchema>;

export const OHLCSchema = z.object({
  date: z.string(),
  open: z.number(),
  high: z.number(),
  low: z.number(),
  close: z.number(),
  volume: z.number(),
});

export type OHLC = z.infer<typeof OHLCSchema>;
