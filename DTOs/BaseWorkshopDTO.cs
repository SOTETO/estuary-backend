using System;

namespace estuary_backend.DTOs
{
    public class BaseWorkshopDTO
    {
        public int id { get; set; }

        public int type { get; set; }

        public string locationName { get; set; }

        public string locationMap { get; set; }

        public DateTime date { get; set; }

        public string[] tags { get; set; }

        public int likes { get; set; }

        public string teaser { get; set; }

    }
}
