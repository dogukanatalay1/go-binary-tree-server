package test

var MaxPathSumTestCases = []struct {
	Name           string
	Input          string
	ExpectedOutput int
}{
	{
		Name: "TestCase",
		Input: `{
			"tree": {
				"nodes": [
					{"id": "1", "left": "2", "right": "3", "value": 1},
					{"id": "3", "left": null, "right": null, "value": 3},
					{"id": "2", "left": null, "right": null, "value": 2}
				],
				"root": "1"
			}
		}`,
		ExpectedOutput: 6,
	},
	{
		Name: "TestCase2",
		Input: `{
			"tree": {
				"nodes": [
					{"id": "1", "left": "-10", "right": "-5", "value": 1},
					{"id": "-5", "left": "-20", "right": "-21", "value": -5},
					{"id": "-21", "left": "100-2", "right": "1-3", "value": -21},
					{"id": "1-3", "left": null, "right": null, "value": 1},
					{"id": "100-2", "left": null, "right": null, "value": 100},
					{"id": "-20", "left": "100", "right": "2", "value": -20},
					{"id": "2", "left": null, "right": null, "value": 2},
					{"id": "100", "left": null, "right": null, "value": 100},
					{"id": "-10", "left": "30", "right": "45", "value": -10},
					{"id": "45", "left": "3", "right": "-3", "value": 45},
					{"id": "-3", "left": null, "right": null, "value": -3},
					{"id": "3", "left": null, "right": null, "value": 3},
					{"id": "30", "left": "5", "right": "1-2", "value": 30},
					{"id": "1-2", "left": null, "right": null, "value": 1},
					{"id": "5", "left": null, "right": null, "value": 5}
				],
				"root": "1"
			}
		}`,
		ExpectedOutput: 154,
	},
}