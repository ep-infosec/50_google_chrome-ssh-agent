load("@bazel_gazelle//:def.bzl", "gazelle")
load("@io_bazel_rules_go//go:def.bzl", "TOOLS_NOGO", "nogo")
load("@rules_pkg//:pkg.bzl", "pkg_zip")

# gazelle:prefix github.com/google/chrome-ssh-agent

# Force Gazelle to choose the correct target when there are multiple go_library
# targets in a single package.
# gazelle:resolve go github.com/google/chrome-ssh-agent/go/agentport //go/agentport
# gazelle:resolve go github.com/google/chrome-ssh-agent/go/chrome //go/chrome
# gazelle:resolve go github.com/google/chrome-ssh-agent/go/dom //go/dom
# gazelle:resolve go github.com/google/chrome-ssh-agent/go/jsutil //go/jsutil
# gazelle:resolve go github.com/google/chrome-ssh-agent/go/keys //go/keys
# gazelle:resolve go github.com/google/chrome-ssh-agent/go/message //go/message
# gazelle:resolve go github.com/google/chrome-ssh-agent/go/message/fakes //go/message/fakes
# gazelle:resolve go github.com/google/chrome-ssh-agent/go/optionsui //go/optionsui
# gazelle:resolve go github.com/google/chrome-ssh-agent/go/storage //go/storage
# gazelle:resolve go github.com/google/chrome-ssh-agent/go/storage/testing //go/storage/testing
# gazelle:resolve go github.com/google/chrome-ssh-agent/go/testutil //go/testutil

gazelle(
    name = "gazelle",
    command = "fix",
)

gazelle(
    name = "gazelle-update-repos",
    args = [
        "--prune=true",
        "--from_file=go.mod",
    ],
    command = "update-repos",
)

# Enable nogo for source code analysis
nogo(
    name = "chrome_ssh_agent_nogo",
    config = ":nogo-config.json",
    visibility = ["//visibility:public"],
    deps = TOOLS_NOGO + [
        # Prohibit build tags (e.g., //go:build ... ) within tests, as these may
        # cause tests to silently be skipped unless we are extremely careful.
        "//nogo/testbuildtags",
    ],
)

exports_files([
    "go.mod",
    "go.sum",
    "package.json",
    "package-lock.json",
])

pkg_zip(
    name = "chrome-ssh-agent",
    srcs = [
        ":CONTRIBUTING.md",
        ":LICENSE",
        ":README.md",
        ":manifest.json",
        "//go/background:pkg",
        "//go/options:pkg",
        "//html:pkg",
        "//img:pkg",
    ],
    visibility = ["//visibility:public"],
)
