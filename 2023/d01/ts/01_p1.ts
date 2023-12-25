import assert from 'node:assert';
import fs from 'node:fs';

function getCalibrationValue(symbols: string): number {
  const arr = symbols.split('');
  const firstNum = (chars: string[]) => chars.find((c) => !isNaN(Number(c))) || '';
  return Number(firstNum(arr) + firstNum(arr.reverse()));
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
const inputs = ['1abc2', 'pqr3stu8vwx', 'a1b2c3d4e5f', 'treb7uchet'];
const answers = [12, 38, 15, 77];

answers.every((answer, index) => {
  assert.strictEqual(getCalibrationValue(inputs[index]), answer);
  return true;
});

assert.strictEqual(solve(inputs), 142);
