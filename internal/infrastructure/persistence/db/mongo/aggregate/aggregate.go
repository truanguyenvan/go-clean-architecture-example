package aggregate

import "go.mongodb.org/mongo-driver/bson"

type Operation bson.M

// The Pipe functions similar to the RxJS pipe function
// (https://rxjs-dev.firebaseapp.com/api/index/function/pipe) in that
// it accepts a set of functions as parameters. Each function receives
// as its input, the output of the previous function. The allows you
// to chain together functions that have the same bson.A signature
// and create elaborate aggregate pipelines that get injected into
// the MongoDB driver. The Pipe function accepts an initial state as
// its first parameter.
func Pipe(pipeline bson.A, steps ...func(bson.A) bson.A) bson.A {
	for _, step := range steps {
		// Iterate through each step and call each step passing in
		// the current state of the pipeline. Set the new state of
		// the pipeline equal the mutated state from the step
		pipeline = step(pipeline)
	}

	// Return the finallized pipeline after completing all the steps
	// to the caller. Usually this is what gets provided to the MongoDB
	// collection.Aggregate function
	return pipeline
}

// The following operators can be provided to the Pipe
// function as parameters. The operators have a special signature that
// allow you to daisy-chain them together.

// The Match function builds the $match operator into a pipeline
func Match(fields Operation) func(bson.A) (a bson.A) {
	// Every operator returns a function which accepts the current state
	// of the pipeline as a parameter. The Aggregate Pipe makes sure that
	// the pipeline as returned from the previous operator is passed into
	// this function in the next operator.

	// This function returns the modified state of the pipeline and is
	// what gets piped into the next operator.
	return func(pipeline bson.A) bson.A {
		// We want to take the current state of the pipeline and append
		// the new raw bson.M, MongoDB operator into the pipeline
		return append(pipeline, bson.M{
			"$match": fields,
		})
	}
}

// The Unwind function builds the $unwind operator into a pipeline
func Unwind(fields Operation) func(bson.A) bson.A {
	return func(pipeline bson.A) bson.A {
		return append(pipeline, bson.M{
			"$unwind": fields,
		})
	}
}

// The UnwindSingle function builds the $unwind operator into a pipeline
func UnwindSingle(field string) func(bson.A) bson.A {
	return func(pipeline bson.A) bson.A {
		return append(pipeline, bson.M{
			"$unwind": field,
		})
	}
}

// The Project function builds the $project operator into a pipeline
func Project(fields Operation) func(bson.A) bson.A {
	return func(pipeline bson.A) bson.A {
		return append(pipeline, bson.M{
			"$project": fields,
		})
	}
}

// The Project function builds the $project operator into a pipeline
func ReplaceRoot(fields Operation) func(bson.A) bson.A {
	return func(pipeline bson.A) bson.A {
		return append(pipeline, bson.M{
			"$replaceRoot": fields,
		})
	}
}

// The Sort function builds the $sort operator into a pipeline
func Sort(fields Operation) func(bson.A) bson.A {
	return func(pipeline bson.A) bson.A {
		return append(pipeline, bson.M{
			"$sort": fields,
		})
	}
}

// The Lookup function builds the $lookup operator into a pipeline
func Lookup(from string, localField string, foreignField string, as string) func(bson.A) bson.A {
	return func(pipeline bson.A) bson.A {
		return append(pipeline, bson.M{
			"$lookup": bson.D{
				{"from", from},
				{"localField", localField},
				{"foreignField", foreignField},
				{"as", as},
			},
		})
	}
}

// The Limit function builds the $limit operator into a pipeline
func Limit(num int) func(bson.A) bson.A {
	return func(pipeline bson.A) bson.A {
		return append(pipeline, bson.M{
			"$limit": num,
		})
	}
}

// The Skip function builds the $skip operator into a pipeline
func Skip(num int) func(bson.A) bson.A {
	return func(pipeline bson.A) bson.A {
		return append(pipeline, bson.M{
			"$skip": num,
		})
	}
}
