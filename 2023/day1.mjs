import { open } from "node:fs/promises"

const readFile = async () => {
	const file = await open("./Day1Input.txt");

	let sum = 0;

	for await (const line of file.readLines()) {
		let numsOnly = line.replace(/[a-zA-z]/g, "");
		console.log(numsOnly[0], numsOnly[numsOnly.length - 1])
		const localSum = parseInt(numsOnly[0]) + parseInt(numsOnly[numsOnly.length - 1])
		sum += localSum
	}

	console.log(sum)
}

readFile();