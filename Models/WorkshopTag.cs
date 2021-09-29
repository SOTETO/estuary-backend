using System;

namespace estuary_backend.Models
{
    public class WorkshopTag
    {
        public Guid Id { get; set; }
        public Guid WorkshopId { get; set; }
        public Workshop Workshop { get; set; }
        public Guid TagId { get; set; }
        public Tag Tag { get; set; }
    }
}
