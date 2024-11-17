// Author: fluffelpuff
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package cenvxcore

type ApiError struct {
	ResponseMessageText string
}

type ExtModCGOPanic struct {
	ErrorValue error
}

type ExtModFunctionCallError struct {
	ErrorValue error
}

type SpecificError struct {
	CliError            error
	LocalJSVMError      error
	GoProcessError      error
	LocalApiOrRpcError  ApiError
	RemoteApiOrRpcError ApiError
	History             []string
}

func (e *ExtModCGOPanic) Error() string {
	return e.ErrorValue.Error()
}

func (e *ExtModFunctionCallError) Error() string {
	return e.ErrorValue.Error()
}

func (e *SpecificError) Error() string {
	return e.GoProcessError.Error()
}

func (e *SpecificError) GetRemoteApiOrRpcErrorMessage() string {
	return ""
}

func (e *SpecificError) GetLocalApiOrRpcErrorMessage() string {
	return ""
}

func (e *SpecificError) GetGoProcessErrorMessage() string {
	return ""
}

func (e *SpecificError) AddCallerFunctionToHistory(funcName string) *SpecificError {
	return nil
}
