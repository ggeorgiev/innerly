load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "com_github_google_uuid",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/google/uuid",
        sum = "h1:KjJaJ9iWZ3jOFZIf1Lqf4laDRCasjl0BCmnEGxkdLb4=",
        version = "v1.3.1",
    )

    go_repository(
        name = "com_github_neo4j_neo4j_go_driver_v5",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/neo4j/neo4j-go-driver/v5",
        sum = "h1:NmyUxh4LYTdcJdI6EnazHyUKu1f0/BPiHCYUZUZIGQw=",
        version = "v5.13.0",
    )
