Name:    python3-provider
Version: 3.8
Release: 1
Summary: Provide python3 via python38 alternatives
License: Apache-2.0
BuildArch: noarch
Requires: python38
Provides: python3 = 3.8
Conflicts: python36

%description
Sets python3 to python3.8 via alternatives.

%post
alternatives --set python3 /usr/bin/python3.8

%files
