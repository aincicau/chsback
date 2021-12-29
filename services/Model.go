package services

import (
	"fmt"

	tf "github.com/galeone/tensorflow/tensorflow/go"
	tg "github.com/galeone/tfgo"
)

var model *tg.Model

func LoadModel() {
	model := tg.LoadModel("model/output/keras", []string{"serve"}, nil)
	fakeInput, _ := tf.NewTensor([1][28][28][1]float32{})
	results := model.Exec([]tf.Output{
		model.Op("StatefulPartitionedCall", 0),
	}, map[tf.Output]*tf.Tensor{
		model.Op("serving_default_inputs_input", 0): fakeInput,
	})
	predictions := results[0]
	fmt.Println(predictions.Value())
}
