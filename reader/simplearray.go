package reader;

import (
    "github.com/wenkesj/rphash/decoder"
    "github.com/wenkesj/rphash/types"
    "github.com/wenkesj/rphash/utils"
);

type SimpleArray struct {
    data [][]float64
    dimension int;
    numberOfProjections int;
    decoderMultiplier int;
    randomSeed int64;
    hashModulus int32;
    k int;
    numberOfBlurs int;
    decoder types.Decoder;
    centroids [][]float64;
    topIDs []int32;
};

func NewSimpleArray(X [][]float64, k int) *SimpleArray {
    var randomSeed int64 = 0;
    innerDecoder := decoder.InnerDecoder();
    hashModulus := int32(2147483647);
    decoderMultiplier := 1;
    decoder := decoder.NewMultiDecoder(decoderMultiplier * innerDecoder.GetDimensionality(), innerDecoder);
    numberOfProjections := 2;
    numberOfBlurs := 2;
    data := X;
    dimension := 0;
    if data != nil {
        dimension = len(data[0]);
    } else {
        dimension = 0;
    }
    centroids := [][]float64{};
    topIDs := []int32{};
    return &SimpleArray{
        data: data,
        dimension: dimension,
        numberOfProjections: numberOfProjections,
        decoderMultiplier: decoderMultiplier,
        randomSeed: randomSeed,
        hashModulus: hashModulus,
        k: k,
        numberOfBlurs: numberOfBlurs,
        decoder: decoder,
        centroids: centroids,
        topIDs: topIDs,
    };
};

func (this *SimpleArray) GetVectorIterator() [][]float64 {
	return this.data;
};

func (this *SimpleArray) GetK() int {
	return this.k;
};

func (this *SimpleArray) GetDimension() int {
	if this.dimension == 0 {
        this.dimension = len(this.data[0]);
    }
	return this.dimension;
};

func (this *SimpleArray) GetHashModulus() int32 {
	return this.hashModulus;
};

func (this *SimpleArray) GetRandomSeed() int64 {
	return this.randomSeed;
};

func (this *SimpleArray) AddCentroid(v []float64) {
	this.centroids = append(this.centroids, v);
};

func (this *SimpleArray) SetCentroids(l [][]float64) {
	this.centroids = l;
};

func (this *SimpleArray) GetCentroids() [][]float64 {
	return this.centroids;
};

func (this *SimpleArray) GetNumberOfBlurs() int {
	return this.numberOfBlurs;
};

func (this *SimpleArray) GetPreviousTopID() []int32 {
	return this.topIDs;
};

func (this *SimpleArray) SetPreviousTopID(top []int32) {
	this.topIDs = top;
};

func (this *SimpleArray) SetRandomSeed(parseLong int64) {
    this.randomSeed = parseLong;
};

func (this *SimpleArray) SetNumberOfProjections(probes int) {
	this.numberOfProjections = probes;
};

func (this *SimpleArray) GetNumberOfProjections() int {
	return this.numberOfProjections;
};

func (this *SimpleArray) SetInnerDecoderMultiplier(multiDim int) {
	this.decoderMultiplier = multiDim;
};

func (this *SimpleArray) GetInnerDecoderMultiplier() int {
	return this.decoderMultiplier;
};

func (this *SimpleArray) SetHashModulus(parseLong int32) {
	this.hashModulus = parseLong;
};

func (this *SimpleArray) SetDecoderType(decoder types.Decoder) {
	this.decoder = decoder;
};

func (this *SimpleArray) SetVariance(data [][]float64) {
	this.decoder.SetVariance(utils.VarianceSample(data, 0.01));
};

func (this *SimpleArray) GetDecoderType() types.Decoder {
	return this.decoder;
};
