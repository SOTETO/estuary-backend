using System;
using Microsoft.EntityFrameworkCore.Migrations;

namespace estuary_backend.Migrations
{
    public partial class InitialCreate : Migration
    {
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.CreateTable(
                name: "Workshops",
                columns: table => new
                {
                    Id = table.Column<Guid>(type: "TEXT", nullable: false),
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
                    ID = table.Column<string>(type: "TEXT", nullable: false),
                    Visible = table.Column<bool>(type: "INTEGER", nullable: false),
                    WorkshopId = table.Column<Guid>(type: "TEXT", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Author", x => x.ID);
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
                    Id = table.Column<Guid>(type: "TEXT", nullable: false),
                    Title = table.Column<string>(type: "TEXT", nullable: true),
                    WorkshopId = table.Column<Guid>(type: "TEXT", nullable: false),
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
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "Tags",
                columns: table => new
                {
                    Id = table.Column<Guid>(type: "TEXT", nullable: false),
                    Name = table.Column<string>(type: "TEXT", nullable: true),
                    WorkshopId = table.Column<Guid>(type: "TEXT", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Tags", x => x.Id);
                    table.ForeignKey(
                        name: "FK_Tags_Workshops_WorkshopId",
                        column: x => x.WorkshopId,
                        principalTable: "Workshops",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Restrict);
                });

            migrationBuilder.CreateTable(
                name: "ContentLink",
                columns: table => new
                {
                    Id = table.Column<Guid>(type: "TEXT", nullable: false),
                    ContentId = table.Column<Guid>(type: "TEXT", nullable: true),
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
                    Id = table.Column<Guid>(type: "TEXT", nullable: false),
                    UserId = table.Column<string>(type: "TEXT", nullable: true),
                    ContentId = table.Column<Guid>(type: "TEXT", nullable: true)
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

            migrationBuilder.CreateTable(
                name: "WorkshopTag",
                columns: table => new
                {
                    Id = table.Column<Guid>(type: "TEXT", nullable: false),
                    WorkshopId = table.Column<Guid>(type: "TEXT", nullable: false),
                    TagId = table.Column<Guid>(type: "TEXT", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_WorkshopTag", x => x.Id);
                    table.ForeignKey(
                        name: "FK_WorkshopTag_Tags_TagId",
                        column: x => x.TagId,
                        principalTable: "Tags",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                    table.ForeignKey(
                        name: "FK_WorkshopTag_Workshops_WorkshopId",
                        column: x => x.WorkshopId,
                        principalTable: "Workshops",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.InsertData(
                table: "Tags",
                columns: new[] { "Id", "Name", "WorkshopId" },
                values: new object[] { new Guid("95a53535-2633-41a5-802b-7ad57efcc0ef"), "Improvements", null });

            migrationBuilder.InsertData(
                table: "Workshops",
                columns: new[] { "Id", "Date", "LocationMap", "LocationName", "Teaser" },
                values: new object[] { new Guid("17bb47aa-c617-40c7-b88d-a3fed6bdc03e"), new DateTime(2021, 9, 29, 17, 41, 41, 97, DateTimeKind.Local).AddTicks(483), "https://goo.gl/maps/mbnen1jr8C81J6vU9", "Hamburg", "Lorem ipsum dolor sit amet." });

            migrationBuilder.InsertData(
                table: "Content",
                columns: new[] { "Id", "Because", "But", "Discriminator", "Feel", "Iam", "Iwant", "Title", "WorkshopId" },
                values: new object[] { new Guid("2cad981e-1de5-4d73-8b3c-b39de97bfe99"), "sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum.", "sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat", "PropblemStatement", "Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.Lorem ipsum dolor sit amet, consetetur sadipscing.", "Lorem ipsum", "dolor sit amet, consetetur sadipscing elitr", "ProblemStatement 1", new Guid("17bb47aa-c617-40c7-b88d-a3fed6bdc03e") });

            migrationBuilder.InsertData(
                table: "WorkshopTag",
                columns: new[] { "Id", "TagId", "WorkshopId" },
                values: new object[] { new Guid("facd3370-3cee-4af7-94a1-5f16a4769e64"), new Guid("95a53535-2633-41a5-802b-7ad57efcc0ef"), new Guid("17bb47aa-c617-40c7-b88d-a3fed6bdc03e") });

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
                name: "IX_Tags_WorkshopId",
                table: "Tags",
                column: "WorkshopId");

            migrationBuilder.CreateIndex(
                name: "IX_WorkshopTag_TagId",
                table: "WorkshopTag",
                column: "TagId");

            migrationBuilder.CreateIndex(
                name: "IX_WorkshopTag_WorkshopId",
                table: "WorkshopTag",
                column: "WorkshopId");
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
                name: "WorkshopTag");

            migrationBuilder.DropTable(
                name: "Content");

            migrationBuilder.DropTable(
                name: "Tags");

            migrationBuilder.DropTable(
                name: "Workshops");
        }
    }
}
