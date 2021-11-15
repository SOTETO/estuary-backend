namespace estuary_backend.DTOs
{
    public class ProblemStatementDTO : WorkshopContentDTO
    {
        public int id { get; set; }

        public int likes { get; set; }

        public string iAm { get; set; }

        public string iWant { get; set; }

        public string but { get; set; }

        public string because { get; set; }

        public string feel { get; set; }

        public ContentLinkDTO[] linked { get; set; }
    }
}
