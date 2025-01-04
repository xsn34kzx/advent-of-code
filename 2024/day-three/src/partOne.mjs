import { readFile } from "fs";

if (process.argv.length !== 2) {
    const filePath = process.argv[2]

    readFile(filePath, { encoding: "utf-8" }, (err, data) => {
        if (err) {
            console.log(err)
        } else {
            const instrRegEx = /(mul\([0-9]+,[0-9]+\))/g
            const instructions = data.match(instrRegEx)
            let sum = 0

            for (const instruction of instructions) {
                const numRegEx = /[0-9]+/g
                const nums = instruction.match(numRegEx)

                const firstVal = Number.parseInt(nums[0])
                const secondVal = Number.parseInt(nums[1])

                sum += firstVal * secondVal 
            }

            console.log(sum)
        }
    });
} else {
    console.log("Must provide two arguments to program!")
}

