using System;
using Microsoft.EntityFrameworkCore.Migrations;

namespace estuary_backend.Migrations
{
    public partial class Init : Migration
    {
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.CreateTable(
                name: "Tags",
                columns: table => new
                {
                    Id = table.Column<int>(type: "INTEGER", nullable: false)
                        .Annotation("Sqlite:Autoincrement", true),
                    Name = table.Column<string>(type: "TEXT", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Tags", x => x.Id);
                });

            migrationBuilder.CreateTable(
                name: "Workshops",
                columns: table => new
                {
                    Id = table.Column<int>(type: "INTEGER", nullable: false)
                        .Annotation("Sqlite:Autoincrement", true),
                    Date = table.Column<DateTime>(type: "TEXT", nullable: false),
                    Teaser = table.Column<string>(type: "TEXT", nullable: true),
                    LocationName = table.Column<string>(type: "TEXT", nullable: true),
                    LocationMap = table.Column<string>(type: "TEXT", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Workshops", x => x.Id);
                });

            migrationBuilder.CreateTable(
                name: "Author",
                columns: table => new
                {
                    Id = table.Column<int>(type: "INTEGER", nullable: false)
                        .Annotation("Sqlite:Autoincrement", true),
                    DropsID = table.Column<string>(type: "TEXT", nullable: true),
                    Visible = table.Column<bool>(type: "INTEGER", nullable: false),
                    WorkshopId = table.Column<int>(type: "INTEGER", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Author", x => x.Id);
                    table.ForeignKey(
                        name: "FK_Author_Workshops_WorkshopId",
                        column: x => x.WorkshopId,
                        principalTable: "Workshops",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Restrict);
                });

            migrationBuilder.CreateTable(
                name: "Content",
                columns: table => new
                {
                    Id = table.Column<int>(type: "INTEGER", nullable: false)
                        .Annotation("Sqlite:Autoincrement", true),
                    Title = table.Column<string>(type: "TEXT", nullable: true),
                    WorkshopId = table.Column<int>(type: "INTEGER", nullable: true),
                    Discriminator = table.Column<string>(type: "TEXT", nullable: false),
                    Iam = table.Column<string>(type: "TEXT", nullable: true),
                    Iwant = table.Column<string>(type: "TEXT", nullable: true),
                    But = table.Column<string>(type: "TEXT", nullable: true),
                    Because = table.Column<string>(type: "TEXT", nullable: true),
                    Feel = table.Column<string>(type: "TEXT", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Content", x => x.Id);
                    table.ForeignKey(
                        name: "FK_Content_Workshops_WorkshopId",
                        column: x => x.WorkshopId,
                        principalTable: "Workshops",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Restrict);
                });

            migrationBuilder.CreateTable(
                name: "TagWorkshop",
                columns: table => new
                {
                    TagsId = table.Column<int>(type: "INTEGER", nullable: false),
                    WorkshopsId = table.Column<int>(type: "INTEGER", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_TagWorkshop", x => new { x.TagsId, x.WorkshopsId });
                    table.ForeignKey(
                        name: "FK_TagWorkshop_Tags_TagsId",
                        column: x => x.TagsId,
                        principalTable: "Tags",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                    table.ForeignKey(
                        name: "FK_TagWorkshop_Workshops_WorkshopsId",
                        column: x => x.WorkshopsId,
                        principalTable: "Workshops",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "ContentLink",
                columns: table => new
                {
                    Id = table.Column<int>(type: "INTEGER", nullable: false)
                        .Annotation("Sqlite:Autoincrement", true),
                    ContentId = table.Column<int>(type: "INTEGER", nullable: true),
                    LinkTag = table.Column<string>(type: "TEXT", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_ContentLink", x => x.Id);
                    table.ForeignKey(
                        name: "FK_ContentLink_Content_ContentId",
                        column: x => x.ContentId,
                        principalTable: "Content",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Restrict);
                });

            migrationBuilder.CreateTable(
                name: "Like",
                columns: table => new
                {
                    Id = table.Column<int>(type: "INTEGER", nullable: false)
                        .Annotation("Sqlite:Autoincrement", true),
                    UserId = table.Column<string>(type: "TEXT", nullable: true),
                    ContentId = table.Column<int>(type: "INTEGER", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Like", x => x.Id);
                    table.ForeignKey(
                        name: "FK_Like_Content_ContentId",
                        column: x => x.ContentId,
                        principalTable: "Content",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Restrict);
                });

            migrationBuilder.CreateIndex(
                name: "IX_Author_WorkshopId",
                table: "Author",
                column: "WorkshopId");

            migrationBuilder.CreateIndex(
                name: "IX_Content_WorkshopId",
                table: "Content",
                column: "WorkshopId");

            migrationBuilder.CreateIndex(
                name: "IX_ContentLink_ContentId",
                table: "ContentLink",
                column: "ContentId");

            migrationBuilder.CreateIndex(
                name: "IX_Like_ContentId",
                table: "Like",
                column: "ContentId");

            migrationBuilder.CreateIndex(
                name: "IX_TagWorkshop_WorkshopsId",
                table: "TagWorkshop",
                column: "WorkshopsId");
        }

        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropTable(
                name: "Author");

            migrationBuilder.DropTable(
                name: "ContentLink");

            migrationBuilder.DropTable(
                name: "Like");

            migrationBuilder.DropTable(
                name: "TagWorkshop");

            migrationBuilder.DropTable(
                name: "Content");

            migrationBuilder.DropTable(
                name: "Tags");

            migrationBuilder.DropTable(
                name: "Workshops");
        }
    }
}
