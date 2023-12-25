import assert from 'node:assert';
import fs from 'node:fs';

const rev = (s: string) => s.split('').reverse().join('');

const names = ['zero', 'one', 'two', 'three', 'four', 'five', 'six', 'seven', 'eight', 'nine'];
const forward = new RegExp(names.join('|') + '|[0-9]');
const backward = new RegExp(names.map(rev).join('|') + '|[0-9]');

function toNumber(str: string): number {
  const num = Number(str);
  return isNaN(num) ? names.indexOf(str) : num;
}

function getCalibrationValue(symbols: string): number {
  if (!symbols) return 0;
  const a = symbols.match(forward)?.[0] || '';
  // find the last one with another reversed regexp to avoid missing overlapping like `twone`
  const b = rev(rev(symbols).match(backward)?.[0] || '');
  return +`${toNumber(a)}${toNumber(b)}`;
}

function solve(inputs: string[]): number {
  return inputs.reduce((acc, input) => acc + getCalibrationValue(input), 0);
}

const inputFile = process.argv[2];

if (inputFile) {
  const file = fs.readFileSync(inputFile, 'utf8');
  const inputs = file.split('\n');
  console.log(solve(inputs));
  process.exit(0);
}

// tests
const inputs = [
  'two1nine',
  'eightwothree',
  'abcone2threexyz',
  'xtwone3four',
  '4nineeightseven2',
  'zoneight234',
  '7pqrstsixteen',
  '5onetwoneg',
];
const answers = [29, 83, 13, 24, 42, 14, 76, 51];

answers.every((answer, index) => {
  assert.strictEqual(getCalibrationValue(inputs[index]), answer);
  return true;
});

assert.strictEqual(solve(inputs), 332);
