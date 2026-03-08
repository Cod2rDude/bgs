package accumulator

/*
    ******************************************************************************
    * @file     : accumulator/accumulator.go
    * @author   : Cod2rDude
    * @date     : March 8 2026
    * @lastEdit : March 8 2026 @ 07:31
    * @brief    : Accumulator is an object that accumulates a value over time.
    * @version  : 1.1.5
    ******************************************************************************
    * @attention
    *
    * Copyright © 2026 Axon Corporation.
    * All rights reserved.
    *
    * This software is licensed under terms that can be found in the LICENSE file
    * in the root directory of this software component.
    * If no LICENSE file comes with this software, it is provided AS-IS.
    *
    ******************************************************************************
*/

// Public Types

/*
    @object Accumulator
    
    @publicVariables
    *   @publicVariable Value uint64 ;; Current value of the accumulator.
    *   @publicVariable InitialValue uint64 ;; Value of accumator when reseted or initiated.
    *   @publicVariable LowerBound uint64 ;; Minimum value that the accumulator can have.
    *   @publicVariable UpperBound uint64 ;; Maximum value that the accumulator can have.
    *   @publicVariable Incrementation uint64 ;; Amount to add to the accumulator each time @publicMethod Accumulate is called.
    *   @publicVariable ResetWhenReachedLimit bool ;; Whatever to if reset value to @publicVariable InitialValue when upper bound is reached.
    *   @publicVariable DoWhenReachedLimit func() ;; Function to call when upper bound is reached or passed. (This can be called more than once.)
    @publicMethods
    *   @publicMethod Accumulate    : () -> ()
    @brief Accumulator is a utility struct that helps to accumulate a value over time, 
    *       with specified bounds and actions when limits are reached.
*/
type Accumulator struct {
    Value                   uint64
    InitialValue            uint64
    LowerBound              uint64
    UpperBound              uint64
    Incrementation          uint64
    ResetWhenReachedLimit   bool
    DoWhenReachedLimit      func()
}

// Public Constructors

/*
    @constructor Accumulator

    @params
    *   @param initialValue uint64 ;; Initial value of the accumulator. If it's out of bounds, it will be clamped to the nearest bound.
    *   @param lowerBound uint64 ;; Minimum value that the accumulator can have. If it's greater than upperBound, they will be swapped.
    *   @param upperBound uint64 ;; Maximum value that the accumulator can have. If it's less than lowerBound, they will be swapped.
    *   @param incrementation uint64 ;; Amount to add to the accumulator each time @publicMethod Accumulate is called.
    *   @param resetWhenReachedLimit bool ;; Whatever to if reset value to @publicVariable InitialValue when upper bound is reached.
    *   @param DoWhenReachedLimit func() ;; Function to call when upper bound is reached or passed. (This can be called more than once.)
    @brief Creates a new Accumulator with the specified parameters.
    @returns
    *   @r1 *Accumulator ;; The newly created Accumulator object.
*/
func New(initialValue, lowerBound, upperBound, incrementation uint64,
        resetWhenReachedLimit bool, doWhenReachedLimit func()) *Accumulator {

    if lowerBound > upperBound {
        lowerBound, upperBound = upperBound, lowerBound
    }

    if initialValue < lowerBound {
        initialValue = lowerBound
    } else if initialValue > upperBound {
        initialValue = upperBound
    }

    /*
    if incrementation == 0 {
        incrementation = 1
    }
    */

    return &Accumulator{
        Value:                  initialValue,
        InitialValue:           initialValue,
        LowerBound:             lowerBound,
        UpperBound:             upperBound,
        Incrementation:         incrementation,
        ResetWhenReachedLimit:  resetWhenReachedLimit,
        DoWhenReachedLimit:     doWhenReachedLimit,
    }
}

// Public Methods

/*
    @brief Accumulate adds @publicVariable Incrementation to @publicVariable Value,
    *       and checks if it has reached or passed the upper bound.
    *       If it has, it calls @publicVariable DoWhenReachedLimit (if it's not nil) and resets 
    *       @publicVariable Value to @publicVariable InitialValue if @publicVariable ResetWhenReachedLimit is true.
*/
func (a *Accumulator) Accumulate() {
    if ^uint64(0) - a.Value < a.Incrementation {
        a.Value = ^uint64(0)
    } else {
        a.Value += a.Incrementation
    }

    // Don't try to check if it equals to upper bound because,
    // If incrementation is good enough it can skip upper bound.
    if a.Value < a.UpperBound {
        return
    }

    if a.DoWhenReachedLimit != nil {
        a.DoWhenReachedLimit()
    }
    if a.ResetWhenReachedLimit {
        a.Value = a.InitialValue
    }
}