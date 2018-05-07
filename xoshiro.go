// Package xoshiro implements the xoshiro256** RNG
/*

Translated from

	http://xoshiro.di.unimi.it/xoshiro256starstar.c

	Scrambled Linear Pseudorandom Number Generators
	David Blackman, Sebastiano Vigna
	https://arxiv.org/abs/1805.01407

	http://www.pcg-random.org/posts/a-quick-look-at-xoshiro256.html
*/
package xoshiro

import "math/bits"

// State is the state for the random number generator
type State struct {
	// struct instead of array due to https://github.com/golang/go/issues/15925
	s0, s1, s2, s3 uint64
}

// New returns a new RNG with the state constructed from seed
func New(seed uint64) *State {
	seed64 := splitMix64(seed)
	return &State{seed64.Next(), seed64.Next(), seed64.Next(), seed64.Next()}
}

// NewFromState returns a new RNG with the supplied state
func NewFromState(seed [4]uint64) *State {
	return &State{seed[0], seed[1], seed[2], seed[3]}
}

// Next returns the next integer in the sequence
func (s *State) Next() uint64 {
	r := bits.RotateLeft64(s.s1*5, 7) * 9

	t := s.s1 << 17

	s.s2 ^= s.s0
	s.s3 ^= s.s1
	s.s1 ^= s.s2
	s.s0 ^= s.s3

	s.s2 ^= t

	s.s3 = bits.RotateLeft64(s.s3, 45)

	return r
}

var jump = [...]uint64{0x180ec6d33cfd0aba, 0xd5a61266f0c9392c, 0xa9582618e03fc9aa, 0x39abdc4529b1661c}

// Jump advances the stream 2^128 steps
func (s *State) Jump() {
	var s0, s1, s2, s3 uint64
	for _, v := range jump[:] {
		for b := uint(0); b < 64; b++ {
			if v&(uint64(1)<<b) != 0 {
				s0 ^= s.s0
				s1 ^= s.s1
				s2 ^= s.s2
				s3 ^= s.s3
			}
			s.Next()
		}
	}

	*s = State{s0, s1, s2, s3}
}

type splitMix64 uint64

// Next returns the next number in the sequence
func (x *splitMix64) Next() uint64 {
	*x += 0x9E3779B97F4A7C15
	z := uint64(*x)
	z = (z ^ (z >> 30)) * 0xBF58476D1CE4E5B9
	z = (z ^ (z >> 27)) * 0x94D049BB133111EB
	z = z ^ (z >> 31)
	return z
}
