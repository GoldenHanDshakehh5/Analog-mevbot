/*
	Copyright 2019 whiteblock Inc.
	This file is a part of the genesis.

	Genesis is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	Genesis is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package entity

//Result is the result of executing the command, contains a type and possibly an error
type Result struct {
	//Error is where the error is stored if this result is not a successful result
	Error error
	//Type is the type of result
	Type string
}

//IsSuccess returns whether or not the result indicates success
func (res Result) IsSuccess() bool {
	return res.Error == nil
}

//IsFatal returns true if there is an errr and it is marked as a fatal error,
//meaning it should not be reattempted
func (res Result) IsFatal() bool {
	return res.Error != nil && res.Type == FatalType
}

//IsRequeue returns true if this result indicates that the command should be retried at a
//later pint
func (res Result) IsRequeue() bool {
	return !res.IsSuccess() && !res.IsFatal()
}

const (
	//SuccessType is the type of a successful result
	SuccessType = "Success"
	//TooSoonType is the type of a result from a cmd which tried to execute too soon
	TooSoonType = "TooSoon"
	//FatalType is the type of a result which indicates a fatal error
	FatalType = "Fatal"
	//ErrorType is the generic error type
	ErrorType = "Error"
)

//NewSuccessResult indicates a successful result
func NewSuccessResult() Result {
	return Result{Type: SuccessType, Error: nil}
}

//NewFatalResult creates a fatal error result. Commands with fatal errors are not retried
func NewFatalResult(err error) Result {
	return Result{Type: FatalType, Error: err}
}

//NewErrorResult creates a result which indicates a non-fatal error. Commands with this result should be requeued.
func NewErrorResult(err error) Result {
	return Result{Type: ErrorType, Error: err}
}
