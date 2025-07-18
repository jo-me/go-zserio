package ztype

import (
	zserio "github.com/woven-planet/go-zserio"
)

// OffsetMethod is a function used to set/check bit offsets in the buffer.
type OffsetMethod func(int, int64)

// Array allows representing arrays of any type and serialize them to the zserio format.
type Array[T any, Y IArrayTraits[T]] struct {
	// ArrayTraits are the array traits used.
	ArrayTraits Y

	// The node used by this array for packing
	PackedContext *zserio.PackingContextNode

	// SetOffsetMethod is an optional function to set the offset to the buffer.
	setOffsetMethod   OffsetMethod
	checkOffsetMethod OffsetMethod

	// RawArray is a reference to the raw array.
	RawArray []T

	// FixedSize is the size of the array, if the array is of fixed size
	FixedSize int

	// IsAuto specifies if the array size is automatically calculated.
	IsAuto bool

	// IsPacked specifies if the array is packed.
	IsPacked bool
}

// Size returns the number of elements in an array.
func (array *Array[T, Y]) Size() int {
	return len(array.RawArray)
}

// ZserioBitSize returns the total size of the unpacked array in bits.
func (array *Array[T, Y]) ZserioBitSize(bitPosition int) (int, error) {
	endBitPosition := bitPosition
	size := array.Size()
	if array.IsAuto {
		delta, err := SignedBitSize(int64(size), 4)
		if err != nil {
			return 0, err
		}
		endBitPosition += delta
	}
	if array.ArrayTraits.BitSizeOfIsConstant() && size > 0 {
		var dummy T
		elementSize := array.ArrayTraits.BitSizeOf(dummy, 0)
		if array.setOffsetMethod != nil {
			endBitPosition += size * elementSize
		} else {
			// all elements are spaced in the same way
			endBitPosition = alignTo(8, endBitPosition)
			endBitPosition += elementSize + (size-1)*alignTo(8, elementSize)
		}
	} else {
		for _, element := range array.RawArray {
			if array.setOffsetMethod != nil {
				endBitPosition = alignTo(8, endBitPosition)
			}
			endBitPosition += array.ArrayTraits.BitSizeOf(element, endBitPosition)
		}
	}
	return endBitPosition - bitPosition, nil
}

// BitSizeOfPacked returns the total size of the packed array in bits.
func (array *Array[T, Y]) ZserioBitSizePacked(bitPosition int) (int, error) {
	endBitPosition := bitPosition
	size := array.Size()
	if array.IsAuto {
		delta, err := SignedBitSize(int64(size), 4)
		if err != nil {
			return 0, err
		}
		endBitPosition += delta
	}
	if size > 0 {
		packedTraits := array.ArrayTraits.PackedTraits()

		for _, element := range array.RawArray {
			if array.setOffsetMethod != nil {
				endBitPosition = alignTo(8, endBitPosition)
			}
			delta, err := packedTraits.BitSizeOf(array.PackedContext, endBitPosition, element)
			if err != nil {
				return 0, err
			}
			endBitPosition += delta
		}
	}
	return endBitPosition - bitPosition, nil
}

// Clone does a deep copy of the array.
func (array *Array[T, Y]) Clone() zserio.ZserioType {
	clone := Array[T, Y]{
		ArrayTraits:       array.ArrayTraits,
		RawArray:          array.RawArray,
		IsAuto:            array.IsAuto,
		IsPacked:          array.IsPacked,
		FixedSize:         array.FixedSize,
		PackedContext:     array.PackedContext,
		setOffsetMethod:   array.setOffsetMethod,
		checkOffsetMethod: array.checkOffsetMethod,
	}
	return &clone
}
