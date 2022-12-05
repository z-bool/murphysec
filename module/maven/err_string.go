// Code generated by "stringer -linecomment -type mvnError -output err_string.go"; DO NOT EDIT.

package maven

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ErrMvnDisabled-1]
	_ = x[ErrMvnNotFound-2]
	_ = x[ErrCheckMvnVersion-3]
	_ = x[ErrBadDepsGraph-4]
	_ = x[ErrInvalidCoordinate-5]
	_ = x[ErrArtifactNotFound-6]
	_ = x[ErrGetArtifactFailed-7]
	_ = x[ErrParsePomFailed-8]
	_ = x[ErrOpenProject-9]
	_ = x[ErrPomCircularDependent-10]
	_ = x[ErrBadCoordinate-11]
	_ = x[ErrCouldNotResolve-12]
	_ = x[ErrMvnExitErr-13]
	_ = x[ErrMvnCmd-14]
	_ = x[ErrInspection-15]
}

const _mvnError_name = "maven: mvn command disabledmaven: mvn command not foundmaven: eval mvn version failedmaven: bad dependency graphmaven: invalid coordinatemaven: artifact not foundmaven: get artifact failedmaven: parse pom failedmaven: open project failedmaven: pom file circular dependentmaven: bad coordinatemaven: couldn't resolvemaven: mvn command exit with non-zero codemaven: error during mvn executionmaven: can't inspect the maven project"

var _mvnError_index = [...]uint16{0, 27, 55, 85, 112, 137, 162, 188, 211, 237, 271, 292, 315, 357, 390, 428}

func (i mvnError) String() string {
	i -= 1
	if i < 0 || i >= mvnError(len(_mvnError_index)-1) {
		return "mvnError(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _mvnError_name[_mvnError_index[i]:_mvnError_index[i+1]]
}